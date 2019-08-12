package service

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type grpcService interface {
	Server
	RegisterServer(srv *grpc.Server)
}

// GRPCServer defines the interface for a gRPC server.
type GRPCServer interface {
	Server
}

type grpcServer struct {
	*ServerConfig
	net.Listener
	grpcServer   *grpc.Server
	healthServer *health.Server
}

// NewGRPCServer creates a new listener and gRPC server for a gRPC service.
func NewGRPCServer(config *ServerConfig, grpcSvc grpcService) GRPCServer {
	var creds credentials.TransportCredentials
	var err error
	var lis net.Listener
	var opts []grpc.ServerOption

	if lis, err = net.Listen("tcp", ":"+config.Port); err != nil {
		log.WithFields(log.Fields{
			"err":  err,
			"port": config.Port,
		}).Fatal("Failed to listen on port")
	}

	if creds, err = credentials.NewServerTLSFromFile(config.TLSCert, config.TLSKey); err != nil {
		log.WithField("err", err).Fatal("Failed to generate server credentials")
	}

	opts = append(opts, grpc.Creds(creds))

	logger := log.StandardLogger()
	logrusEntry := log.NewEntry(logger)

	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_logrus.StreamServerInterceptor(logrusEntry),
		streamServerContextInterceptor(logRequestIDFromIncomingContext),
		streamServerContextInterceptor(logClaimsFromIncomingContext),
		streamServerContextInterceptor(logTimingFromIncomingContext),
	))
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_logrus.UnaryServerInterceptor(logrusEntry),
		unaryServerContextInterceptor(logRequestIDFromIncomingContext),
		unaryServerContextInterceptor(logClaimsFromIncomingContext),
		unaryServerContextInterceptor(logTimingFromIncomingContext),
	))

	srv := &grpcServer{config, lis, grpc.NewServer(opts...), health.NewServer()}
	healthpb.RegisterHealthServer(srv.grpcServer, srv.healthServer)
	grpcSvc.RegisterServer(srv.grpcServer)

	return srv
}

func (s *grpcServer) serve() error {
	return s.grpcServer.Serve(s.Listener)
}

func (s *grpcServer) StartServer() {
	s.Serve(s.serve)
}

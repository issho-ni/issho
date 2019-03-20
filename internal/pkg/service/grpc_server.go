package service

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
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
	lis, err := net.Listen("tcp", ":"+config.Port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	creds, err := credentials.NewServerTLSFromFile(config.TLSCert, config.TLSKey)

	if err != nil {
		log.Fatalf("Failed to generate server credentials: %v", err)
	}

	opts = append(opts, grpc.Creds(creds))

	logger := log.StandardLogger()
	logrusEntry := log.NewEntry(logger)

	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_logrus.StreamServerInterceptor(logrusEntry),
		streamServerContextInterceptor(logRequestIDFromIncomingContext),
		streamServerContextInterceptor(logUserIDFromIncomingContext),
	))
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_logrus.UnaryServerInterceptor(logrusEntry),
		unaryServerContextInterceptor(logRequestIDFromIncomingContext),
		unaryServerContextInterceptor(logUserIDFromIncomingContext),
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

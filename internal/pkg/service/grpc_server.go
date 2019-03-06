package service

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type grpcService interface {
	RegisterServer(srv *grpc.Server)
}

type grpcServer struct {
	*ServerConfig
	net.Listener
	grpcServer *grpc.Server
}

// NewGRPCServer creates a new listener and gRPC server for a gRPC service.
func NewGRPCServer(config *ServerConfig, grpcSvc grpcService) Server {
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
		requestIDStreamServerInterceptor))
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_logrus.UnaryServerInterceptor(logrusEntry),
		requestIDUnaryServerInterceptor))

	srv := &grpcServer{config, lis, grpc.NewServer(opts...)}
	grpcSvc.RegisterServer(srv.grpcServer)

	return srv
}

func (s *grpcServer) serve() error {
	return s.grpcServer.Serve(s.Listener)
}

func (s *grpcServer) StartServer() {
	s.Serve(s.serve)
}

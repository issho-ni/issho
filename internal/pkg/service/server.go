package service

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// NewServer creates a new listener and gRPC server for a gRPC service.
func NewServer(port string, tlsCert string, tlsKey string) (net.Listener, *grpc.Server) {
	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	creds, err := credentials.NewServerTLSFromFile(tlsCert, tlsKey)

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

	return lis, grpc.NewServer(opts...)
}

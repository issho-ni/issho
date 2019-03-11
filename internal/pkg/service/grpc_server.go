package service

import (
	"context"
	"net"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type grpcService interface {
	RegisterServer(srv *grpc.Server)
}

type grpcServer struct {
	*ServerConfig
	net.Listener
	mongoClient  *mongo.Client
	grpcServer   *grpc.Server
	healthServer *health.Server
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

	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://issho-mongodb:27017"))
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	srv := &grpcServer{config, lis, mongoClient, grpc.NewServer(opts...), health.NewServer()}
	healthpb.RegisterHealthServer(srv.grpcServer, srv.healthServer)
	grpcSvc.RegisterServer(srv.grpcServer)

	return srv
}

func (s *grpcServer) serve() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := s.mongoClient.Connect(ctx)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	return s.grpcServer.Serve(s.Listener)
}

func (s *grpcServer) StartServer() {
	s.Serve(s.serve)
}

package grpc

import (
	"net"

	"github.com/issho-ni/issho/internal/pkg/mongo"
	"github.com/issho-ni/issho/internal/pkg/service"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Service defines the interface of a server for a gRPC service.
type Service interface {
	service.Server
	RegisterServer(srv *grpc.Server)
}

// Server defines the structure of the server for a gRPC service.
type Server struct {
	*service.ServerConfig
	net.Listener
	MongoClient  mongo.Client
	grpcServer   *grpc.Server
	healthServer *health.Server
}

// NewServer creates a new listener and gRPC server for a gRPC service.
func NewServer(config *service.ServerConfig, srv Service) *Server {
	var creds credentials.TransportCredentials
	var grpcLogrusOpts []grpc_logrus.Option
	var err error
	var opts []grpc.ServerOption
	var s *Server

	s.ServerConfig = config
	s.MongoClient = mongo.NewClient(config.Name)

	if s.Listener, err = net.Listen("tcp", ":"+config.Port); err != nil {
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

	grpcLogrusOpts = append(grpcLogrusOpts, grpc_logrus.WithDecider(logDecider))

	opts = append(opts, grpc_middleware.WithStreamServerChain(
		grpc_logrus.StreamServerInterceptor(logrusEntry, grpcLogrusOpts...),
		streamServerContextInterceptor(logRequestIDFromIncomingContext),
		streamServerContextInterceptor(logClaimsFromIncomingContext),
		streamServerContextInterceptor(logTimingFromIncomingContext),
	))
	opts = append(opts, grpc_middleware.WithUnaryServerChain(
		grpc_logrus.UnaryServerInterceptor(logrusEntry, grpcLogrusOpts...),
		unaryServerContextInterceptor(logRequestIDFromIncomingContext),
		unaryServerContextInterceptor(logClaimsFromIncomingContext),
		unaryServerContextInterceptor(logTimingFromIncomingContext),
	))

	s.grpcServer = grpc.NewServer(opts...)
	s.healthServer = health.NewServer()
	healthpb.RegisterHealthServer(s.grpcServer, s.healthServer)

	srv.RegisterServer(s.grpcServer)
	return s
}

func (s *Server) serve() error {
	return s.grpcServer.Serve(s.Listener)
}

// StartServer provides the callback function to start the server.
func (s *Server) StartServer() {
	cancel := s.MongoClient.Connect()
	defer cancel()

	s.ServerConfig.Serve(s.serve)
}

func logDecider(methodFullName string, err error) bool {
	if err == nil && methodFullName == "/grpc.health.v1.Health/Check" {
		return false
	}

	return true
}

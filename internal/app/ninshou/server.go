package ninshou

import (
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/service"

	"google.golang.org/grpc"
)

type ninshouServer struct {
	service.GRPCServer
	ninshou.NinshouServer
}

func (s *ninshouServer) RegisterServer(srv *grpc.Server) {
	ninshou.RegisterNinshouServer(srv, s)
}

// NewNinshouServer returns a new gRPC server for the Ninshou service.
func NewNinshouServer(config *service.ServerConfig) service.Server {
	server := &ninshouServer{}
	server.GRPCServer = service.NewGRPCServer(config, server)
	return server
}

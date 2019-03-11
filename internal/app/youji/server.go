package youji

import (
	"github.com/issho-ni/issho/api/youji"
	"github.com/issho-ni/issho/internal/pkg/service"

	"google.golang.org/grpc"
)

type youjiServer struct {
	service.GRPCServer
	youji.YoujiServer
}

func (s *youjiServer) RegisterServer(srv *grpc.Server) {
	youji.RegisterYoujiServer(srv, s)
}

// NewYoujiServer returns a new gRPC server for the Youji service.
func NewYoujiServer(config *service.ServerConfig) service.Server {
	server := &youjiServer{}
	server.GRPCServer = service.NewGRPCServer(config, server)
	return server
}

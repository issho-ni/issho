//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I.. ninshou/ninshou.proto

package ninshou

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/service"
)

// Client is the client to the Ninshou gRPC service.
type Client struct {
	service.GRPCClient
	NinshouClient
}

// NewClient returns a client to the Ninshou gRPC service.
func NewClient(config *service.GRPCClientConfig) *Client {
	url := os.Getenv("NINSHOU_URL")
	c := service.NewGRPCClient(config, "ninshou", url)
	return &Client{c, NewNinshouClient(c.ClientConn())}
}

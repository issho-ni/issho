//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. shinninjou/shinninjou.proto

package shinninjou

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/service"
)

// Client is the client to the Shinninjou gRPC service.
type Client struct {
	service.GRPCClient
	ShinninjouClient
}

// NewClient returns a client to the Shinninjou gRPC service.
func NewClient(config *service.GRPCClientConfig) *Client {
	url := os.Getenv("SHINNINJOU_URL")
	c := service.NewGRPCClient(config, "shinninjou", url)
	return &Client{c, NewShinninjouClient(c.ClientConn())}
}

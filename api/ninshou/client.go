//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:. -I../protobuf ninshou.proto

package ninshou

import (
	"github.com/issho-ni/issho/internal/pkg/service"
)

// Client is the client to the Ninshou gRPC service.
type Client struct {
	*service.Client
	NinshouClient
}

// NewClient returns a client to the Ninshou gRPC service.
func NewClient(config *service.ClientConfig) *Client {
	c := service.NewClient(config, "ninshou", "localhost:8081")
	return &Client{c, NewNinshouClient(c.ClientConn)}
}

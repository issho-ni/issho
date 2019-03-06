//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:. -I../protobuf youji.proto

package youji

import (
	"github.com/issho-ni/issho/internal/pkg/service"
)

// Client is the client to the Youji gRPC service.
type Client struct {
	*service.Client
	YoujiClient
}

// NewClient returns a client to the Youji gRPC service.
func NewClient(config *service.ClientConfig) *Client {
	c := service.NewClient(config, "youji", "localhost:8083")
	return &Client{c, NewYoujiClient(c.ClientConn)}
}

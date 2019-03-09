//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I.. youji/youji.proto

package youji

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/service"
)

// Client is the client to the Youji gRPC service.
type Client struct {
	service.Client
	YoujiClient
}

// NewClient returns a client to the Youji gRPC service.
func NewClient(config *service.ClientConfig) *Client {
	url := os.Getenv("YOUJI_URL")
	c := service.NewClient(config, "youji", url)
	return &Client{c, NewYoujiClient(c.ClientConn())}
}

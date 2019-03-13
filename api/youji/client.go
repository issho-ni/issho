//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. youji/youji.proto

package youji

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/service"
)

// Client is the client to the Youji gRPC service.
type Client struct {
	service.GRPCClient
	YoujiClient
}

// NewClient returns a client to the Youji gRPC service.
func NewClient(config *service.GRPCClientConfig) *Client {
	url := os.Getenv("YOUJI_URL")
	c := service.NewGRPCClient(config, "youji", url)
	return &Client{c, NewYoujiClient(c.ClientConn())}
}

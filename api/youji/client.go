//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. youji/youji.proto

package youji

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/grpc"
)

// Client is the client to the Youji gRPC service.
type Client struct {
	grpc.Client
	YoujiClient
}

// NewClient returns a client to the Youji gRPC service.
func NewClient(config *grpc.ClientConfig) *Client {
	url := os.Getenv("YOUJI_URL")
	c := grpc.NewClient(config, "youji", url)
	return &Client{c, NewYoujiClient(c.ClientConn())}
}

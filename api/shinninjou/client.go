//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. shinninjou/shinninjou.proto

package shinninjou

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/grpc"
)

// Client is the client to the Shinninjou gRPC service.
type Client struct {
	grpc.Client
	ShinninjouClient
}

// NewClient returns a client to the Shinninjou gRPC service.
func NewClient(config *grpc.ClientConfig) *Client {
	var client *Client
	url := os.Getenv("SHINNINJOU_URL")
	client.Client = grpc.NewClient(config, "shinninjou", url)
	client.ShinninjouClient = NewShinninjouClient(client.ClientConn())
	return client
}

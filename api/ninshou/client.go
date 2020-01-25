//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. ninshou/ninshou.proto

package ninshou

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/grpc"
)

// Client is the client to the Ninshou gRPC service.
type Client struct {
	grpc.Client
	NinshouClient
}

// NewClient returns a client to the Ninshou gRPC service.
func NewClient(config *grpc.ClientConfig) *Client {
	url := os.Getenv("NINSHOU_URL")
	client := &Client{
		Client: grpc.NewClient(config, "ninshou", url),
	}
	client.NinshouClient = NewNinshouClient(client.ClientConn())
	return client
}

//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. ninka/ninka.proto

package ninka

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/grpc"
)

// Client is the client to the Ninka gRPC service.
type Client struct {
	grpc.Client
	NinkaClient
}

// NewClient returns a client to the Ninka gRPC service.
func NewClient(config *grpc.ClientConfig) *Client {
	url := os.Getenv("NINKA_URL")
	client := &Client{
		Client: grpc.NewClient(config, "ninka", url),
	}
	client.NinkaClient = NewNinkaClient(client.ClientConn())
	return client
}

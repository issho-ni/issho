//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. kazoku/kazoku.proto

package kazoku

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/grpc"
)

// Client is the client for the Kazoku gRPC service.
type Client struct {
	grpc.Client
	KazokuClient
}

// NewClient returns a client to the Kazoku gRPC service.
func NewClient(config *grpc.ClientConfig) *Client {
	var client *Client
	url := os.Getenv("KAZOKU_URL")
	client.Client = grpc.NewClient(config, "kazoku", url)
	client.KazokuClient = NewKazokuClient(client.ClientConn())
	return client
}

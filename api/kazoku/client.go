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
	url := os.Getenv("KAZOKU_URL")
	c := grpc.NewClient(config, "kazoku", url)
	return &Client{c, NewKazokuClient(c.ClientConn())}
}

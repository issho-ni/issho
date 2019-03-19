//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. ninka/ninka.proto

package ninka

import (
	"os"

	"github.com/issho-ni/issho/internal/pkg/service"
)

// Client is the client to the Ninka gRPC service.
type Client struct {
	service.GRPCClient
	NinkaClient
}

// NewClient returns a client to the Ninka gRPC service.
func NewClient(config *service.GRPCClientConfig) *Client {
	url := os.Getenv("NINKA_URL")
	c := service.NewGRPCClient(config, "ninka", url)
	return &Client{c, NewNinkaClient(c.ClientConn())}
}

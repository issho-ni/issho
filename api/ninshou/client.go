//go:generate protoc --gogofaster_out=plugins=grpc:. ninshou.proto

package ninshou

import (
	"github.com/issho-ni/issho/internal/pkg/service"
)

// Client is the client to the Ninshou gRPC service.
type Client struct {
	*service.Client
	NinshouClient
}

// NewClient returns a client to the Ninshou gRPC service.
func NewClient(e service.Env) *Client {
	c := service.NewClient(e, "ninshou", "localhost:8081")
	return &Client{c, NewNinshouClient(c.ClientConn)}
}

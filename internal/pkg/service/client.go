package service

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// Client is the generic client to a gRPC service.
type Client struct {
	*grpc.ClientConn
	e Env
}

// NewClient establishes a client connection to a gRPC service.
func NewClient(e Env, name string, url string) *Client {
	cc, err := grpc.Dial(url, grpc.WithTransportCredentials(e.Creds()))

	if err != nil {
		log.Fatalf("Failed to establish connection to %s: %v", name, err)
	}

	log.Printf("Established connection to %s", name)
	return &Client{cc, e}
}

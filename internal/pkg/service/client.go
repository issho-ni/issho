package service

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
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
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(e.Creds()))
	opts = append(opts, grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
		requestIDStreamClientInterceptor)))
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
		requestIDUnaryClientInterceptor)))

	cc, err := grpc.Dial(url, opts...)

	if err != nil {
		log.Fatalf("Failed to establish connection to %s: %v", name, err)
	}

	log.Printf("Established connection to %s", name)
	return &Client{cc, e}
}

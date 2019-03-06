package service

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// ClientConfig defines the interface for the environment of a service's connection as a
// client to other services.
type ClientConfig struct {
	credentials.TransportCredentials
}

// NewClientConfig generates a new service client environment.
func NewClientConfig(tlsCert string) *ClientConfig {
	creds, err := credentials.NewClientTLSFromFile(tlsCert, "")

	if err != nil {
		log.Fatalf("Failed to generate credentials: %v", err)
	}

	return &ClientConfig{creds}
}

// Client is the generic client to a gRPC service.
type Client struct {
	*grpc.ClientConn
	*ClientConfig
}

// NewClient establishes a client connection to a gRPC service.
func NewClient(config *ClientConfig, name string, url string) *Client {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(config.TransportCredentials))
	opts = append(opts, grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
		requestIDStreamClientInterceptor)))
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
		requestIDUnaryClientInterceptor)))

	cc, err := grpc.Dial(url, opts...)

	if err != nil {
		log.Fatalf("Failed to establish connection to %s: %v", name, err)
	}

	log.Printf("Established connection to %s", name)
	return &Client{cc, config}
}

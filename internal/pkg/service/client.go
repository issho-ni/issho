package service

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
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
type Client interface {
	ClientConn() *grpc.ClientConn
	HealthCheck() bool
}

type client struct {
	cc *grpc.ClientConn
	*ClientConfig
	healthpb.HealthClient
}

// NewClient establishes a client connection to a gRPC service.
func NewClient(config *ClientConfig, name string, url string) Client {
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

	healthClient := healthpb.NewHealthClient(cc)

	log.Printf("Established connection to %s", name)
	return &client{cc, config, healthClient}
}

func (c *client) ClientConn() *grpc.ClientConn {
	return c.cc
}

func (c *client) HealthCheck() bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rpcCtx, rpcCancel := context.WithTimeout(ctx, 1 * time.Second)
	defer rpcCancel()

	resp, err := c.HealthClient.Check(rpcCtx, &healthpb.HealthCheckRequest{})
	if err == nil {
		return resp.GetStatus() == healthpb.HealthCheckResponse_SERVING
	}
	return false
}

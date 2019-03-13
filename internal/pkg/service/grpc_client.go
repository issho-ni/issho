package service

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// GRPCClientConfig defines the interface for the environment of a service's
// connection as a client to other services.
type GRPCClientConfig struct {
	credentials.TransportCredentials
}

// NewGRPCClientConfig generates a new service client environment.
func NewGRPCClientConfig(tlsCert string) *GRPCClientConfig {
	creds, err := credentials.NewClientTLSFromFile(tlsCert, "")

	if err != nil {
		log.Fatalf("Failed to generate credentials: %v", err)
	}

	return &GRPCClientConfig{creds}
}

// GRPCClient is the generic client to a gRPC service.
type GRPCClient interface {
	ClientConn() *grpc.ClientConn
	HealthCheck() bool
}

type grpcClient struct {
	cc *grpc.ClientConn
	*GRPCClientConfig
	healthpb.HealthClient
}

// NewGRPCClient establishes a client connection to a gRPC service.
func NewGRPCClient(config *GRPCClientConfig, name string, url string) GRPCClient {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(config.TransportCredentials))
	opts = append(opts, grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
		requestIDStreamClientInterceptor)))
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
		requestIDUnaryClientInterceptor)))

	cc, err := grpc.Dial(url, opts...)

	if err != nil {
		log.Fatalf("Failed to dial %s: %v", name, err)
	}

	healthClient := healthpb.NewHealthClient(cc)
	log.Debugf("Connecting to %s", name)

	return &grpcClient{cc, config, healthClient}
}

func (c *grpcClient) ClientConn() *grpc.ClientConn {
	return c.cc
}

func (c *grpcClient) HealthCheck() bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rpcCtx, rpcCancel := context.WithTimeout(ctx, 1*time.Second)
	defer rpcCancel()

	resp, err := c.HealthClient.Check(rpcCtx, &healthpb.HealthCheckRequest{})
	if err == nil {
		return resp.GetStatus() == healthpb.HealthCheckResponse_SERVING
	}
	return false
}

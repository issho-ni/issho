package service

import (
	"context"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
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
	var creds credentials.TransportCredentials
	var err error

	if creds, err = credentials.NewClientTLSFromFile(tlsCert, ""); err != nil {
		log.WithField("err", err).Fatal("Failed to generate credentials")
	}

	return &GRPCClientConfig{creds}
}

// GRPCClient is the generic client to a gRPC service.
type GRPCClient interface {
	ClientConn() *grpc.ClientConn
	HealthCheck() *GRPCStatus
}

type grpcClient struct {
	cc *grpc.ClientConn
	*GRPCClientConfig
	healthpb.HealthClient
}

// NewGRPCClient establishes a client connection to a gRPC service.
func NewGRPCClient(config *GRPCClientConfig, name string, url string) GRPCClient {
	var cc *grpc.ClientConn
	var err error
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(config.TransportCredentials))
	opts = append(opts, grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
		streamClientContextInterceptor(appendRequestIDToOutgoingContext),
		streamClientContextInterceptor(appendClaimsToOutgoingContext),
		streamClientContextInterceptor(appendTimingToOutgoingContext),
	)))
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
		unaryClientContextInterceptor(appendRequestIDToOutgoingContext),
		unaryClientContextInterceptor(appendClaimsToOutgoingContext),
		unaryClientContextInterceptor(appendTimingToOutgoingContext),
	)))

	if cc, err = grpc.Dial(url, opts...); err != nil {
		log.WithFields(log.Fields{
			"err":          err,
			"grpc.service": name,
			"span.kind":    "client",
		}).Fatal("Failed to dial")
	}

	healthClient := healthpb.NewHealthClient(cc)

	log.WithFields(log.Fields{
		"grpc.service": name,
		"span.kind":    "client",
	}).Debug("Connecting")
	return &grpcClient{cc, config, healthClient}
}

func (c *grpcClient) ClientConn() *grpc.ClientConn {
	return c.cc
}

// GRPCStatus represents the response to a gRPC health check.
type GRPCStatus struct {
	Result bool
	Error  error
}

func (c *grpcClient) HealthCheck() *GRPCStatus {
	var err error
	var resp *healthpb.HealthCheckResponse

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if resp, err = c.HealthClient.Check(ctx, &healthpb.HealthCheckRequest{}); err != nil {
		return &GRPCStatus{false, err}
	}

	return &GRPCStatus{resp.GetStatus() == healthpb.HealthCheckResponse_SERVING, nil}
}

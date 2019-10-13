package grpc

import (
	"context"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// ClientConfig defines the interface for the environment of a service's
// connection as a client to other services.
type ClientConfig struct {
	credentials.TransportCredentials
}

// NewClientConfig generates a new service client environment.
func NewClientConfig(tlsCert string) *ClientConfig {
	var creds credentials.TransportCredentials
	var err error

	if creds, err = credentials.NewClientTLSFromFile(tlsCert, ""); err != nil {
		log.WithField("err", err).Fatal("Failed to generate credentials")
	}

	return &ClientConfig{creds}
}

// Client is the generic client to a gRPC service.
type Client interface {
	ClientConn() *grpc.ClientConn
	HealthCheck() *Status
}

type client struct {
	cc *grpc.ClientConn
	*ClientConfig
	healthpb.HealthClient
}

// NewClient establishes a client connection to a gRPC service.
func NewClient(config *ClientConfig, name string, url string) Client {
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
	return &client{cc, config, healthClient}
}

func (c *client) ClientConn() *grpc.ClientConn {
	return c.cc
}

// Status represents the response to a gRPC health check.
type Status struct {
	Result bool
	Error  error
}

func (c *client) HealthCheck() *Status {
	var err error
	var resp *healthpb.HealthCheckResponse

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if resp, err = c.HealthClient.Check(ctx, &healthpb.HealthCheckRequest{}); err != nil {
		return &Status{false, err}
	}

	return &Status{resp.GetStatus() == healthpb.HealthCheckResponse_SERVING, nil}
}

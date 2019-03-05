package ninshou

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func requestIDStreamServerInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logRequestIDFromIncomingContext(stream.Context())
	return handler(srv, stream)
}

func requestIDUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logRequestIDFromIncomingContext(ctx)
	return handler(ctx, req)
}

func logRequestIDFromIncomingContext(ctx context.Context) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return
	}

	rid := md.Get("request_id")

	if len(rid) != 1 {
		return
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		"request_id": rid[0],
	})
}

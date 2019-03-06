package service

import (
	"context"

	"github.com/issho-ni/issho/internal/pkg/request_id"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func requestIDStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return streamer(appendRequestIDToOutgoingContext(ctx), desc, cc, method, opts...)
}

func requestIDUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(appendRequestIDToOutgoingContext(ctx), method, req, reply, cc, opts...)
}

func appendRequestIDToOutgoingContext(ctx context.Context) context.Context {
	rid, ok := requestid.FromContext(ctx)

	if ok {
		ctx = metadata.AppendToOutgoingContext(ctx, "request_id", rid.String())
	}

	return ctx
}

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

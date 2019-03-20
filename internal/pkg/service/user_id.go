package service

import (
	"context"

	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func userIDStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return streamer(appendUserIDToOutgoingContext(ctx), desc, cc, method, opts...)
}

func userIDUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(appendUserIDToOutgoingContext(ctx), method, req, reply, cc, opts...)
}

func appendUserIDToOutgoingContext(ctx context.Context) context.Context {
	uid, ok := icontext.UserIDFromContext(ctx)

	if ok {
		ctx = metadata.AppendToOutgoingContext(ctx, "user_id", uid.String())
	}

	return ctx
}

func userIDStreamServerInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logUserIDFromIncomingContext(stream.Context())
	return handler(srv, stream)
}

func userIDUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logUserIDFromIncomingContext(ctx)
	return handler(ctx, req)
}

func logUserIDFromIncomingContext(ctx context.Context) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return
	}

	uid := md.Get("user_id")

	if len(uid) != 1 {
		return
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		"user_id": uid[0],
	})
}

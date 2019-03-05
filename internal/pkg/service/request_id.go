package service

import (
	"context"

	"github.com/issho-ni/issho/internal/pkg/request_id"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func requestIDStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return streamer(appendRequestIDToOutgoingContext(ctx), desc, cc, method, opts...)
}

func requestIDUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	return invoker(appendRequestIDToOutgoingContext(ctx),method, req, reply, cc, opts...)
}

func appendRequestIDToOutgoingContext(ctx context.Context) context.Context {
	rid, ok := requestid.FromContext(ctx)

	if ok {
		ctx = metadata.AppendToOutgoingContext(ctx, "request_id", rid.String())
	}

	return ctx
}

package grpc

import (
	"context"

	"google.golang.org/grpc"
)

func streamClientContextInterceptor(interceptor func(context.Context) context.Context) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		return streamer(interceptor(ctx), desc, cc, method, opts...)
	}
}

func streamServerContextInterceptor(interceptor func(context.Context) context.Context) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		interceptor(ss.Context())
		return handler(srv, ss)
	}
}

func unaryClientContextInterceptor(interceptor func(context.Context) context.Context) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return invoker(interceptor(ctx), method, req, reply, cc, opts...)
	}
}

func unaryServerContextInterceptor(interceptor func(context.Context) context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx = interceptor(ctx)
		return handler(ctx, req)
	}
}

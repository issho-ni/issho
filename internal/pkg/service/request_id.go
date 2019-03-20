package service

import (
	"context"

	icontext "github.com/issho-ni/issho/internal/pkg/context"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func appendRequestIDToOutgoingContext(ctx context.Context) context.Context {
	rid, ok := icontext.RequestIDFromContext(ctx)

	if ok {
		ctx = metadata.AppendToOutgoingContext(ctx, "request_id", rid.String())
	}

	return ctx
}

func logRequestIDFromIncomingContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return ctx
	}

	value := md.Get("request_id")

	if len(value) != 1 {
		return ctx
	}

	rid, err := uuid.Parse(value[0])
	if err != nil {
		return ctx
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		"request_id": rid.String(),
	})

	return icontext.NewRequestIDContext(ctx, rid)
}

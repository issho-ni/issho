package grpc

import (
	"context"

	icontext "github.com/issho-ni/issho/internal/pkg/context"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

const requestIDKey = "request_id"

func appendRequestIDToOutgoingContext(ctx context.Context) context.Context {
	rid, ok := icontext.RequestIDFromContext(ctx)
	if ok {
		rids := rid.String()
		ctx = metadata.AppendToOutgoingContext(ctx, requestIDKey, rids)
	}

	return ctx
}

func logRequestIDFromIncomingContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	value := md.Get(requestIDKey)
	if len(value) != 1 {
		return ctx
	}

	rid, err := uuid.Parse(value[0])
	if err != nil {
		return ctx
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		requestIDKey: rid.String(),
	})

	return icontext.NewRequestIDContext(ctx, rid)
}

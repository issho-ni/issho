package service

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
	var ok bool
	var rid uuid.UUID

	if rid, ok = icontext.RequestIDFromContext(ctx); ok {
		rids := rid.String()
		ctx = metadata.AppendToOutgoingContext(ctx, requestIDKey, rids)
	}

	return ctx
}

func logRequestIDFromIncomingContext(ctx context.Context) context.Context {
	var err error
	var md metadata.MD
	var ok bool
	var rid uuid.UUID
	var value []string

	if md, ok = metadata.FromIncomingContext(ctx); !ok {
		return ctx
	}

	if value = md.Get(requestIDKey); len(value) != 1 {
		return ctx
	}

	if rid, err = uuid.Parse(value[0]); err != nil {
		return ctx
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		requestIDKey: rid.String(),
	})

	return icontext.NewRequestIDContext(ctx, rid)
}

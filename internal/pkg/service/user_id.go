package service

import (
	"context"

	icontext "github.com/issho-ni/issho/internal/pkg/context"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func appendUserIDToOutgoingContext(ctx context.Context) context.Context {
	uid, ok := icontext.UserIDFromContext(ctx)

	if ok {
		ctx = metadata.AppendToOutgoingContext(ctx, "user_id", uid.String())
	}

	return ctx
}

func logUserIDFromIncomingContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return ctx
	}

	value := md.Get("user_id")

	if len(value) != 1 {
		return ctx
	}

	uid, err := uuid.Parse(value[0])
	if err != nil {
		return ctx
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		"user_id": uid.String(),
	})

	return icontext.NewUserIDContext(ctx, uid)
}

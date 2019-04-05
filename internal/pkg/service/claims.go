package service

import (
	"context"
	"encoding/json"

	"github.com/issho-ni/issho/api/common"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func appendClaimsToOutgoingContext(ctx context.Context) context.Context {
	claims, ok := icontext.ClaimsFromContext(ctx)
	if !ok {
		return ctx
	}

	value, err := json.Marshal(claims)
	if err != nil {
		return ctx
	}

	return metadata.AppendToOutgoingContext(ctx, "claims", string(value))
}

func logClaimsFromIncomingContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	value := md.Get("claims")
	if len(value) != 1 {
		return ctx
	}

	var claims common.Claims
	err := json.Unmarshal([]byte(value[0]), &claims)
	if err != nil {
		return ctx
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		"user_id": claims.UserID.String(),
	})

	return icontext.NewClaimsContext(ctx, claims)
}

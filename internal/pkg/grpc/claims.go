package grpc

import (
	"context"
	"encoding/json"

	"github.com/issho-ni/issho/api/common"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

const claimsKey = "claims"
const userIDKey = "user_id"

func appendClaimsToOutgoingContext(ctx context.Context) context.Context {
	var claims common.Claims
	var err error
	var ok bool
	var value []byte

	if claims, ok = icontext.ClaimsFromContext(ctx); !ok {
		return ctx
	}

	if value, err = json.Marshal(claims); err != nil {
		return ctx
	}

	return metadata.AppendToOutgoingContext(ctx, claimsKey, string(value))
}

func logClaimsFromIncomingContext(ctx context.Context) context.Context {
	var claims common.Claims
	var md metadata.MD
	var ok bool
	var value []string

	if md, ok = metadata.FromIncomingContext(ctx); !ok {
		return ctx
	}

	if value = md.Get(claimsKey); len(value) != 1 {
		return ctx
	}

	if err := json.Unmarshal([]byte(value[0]), &claims); err != nil {
		return ctx
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		userIDKey: claims.UserID.String(),
	})

	return icontext.NewClaimsContext(ctx, claims)
}

package context

import (
	"context"

	"github.com/issho-ni/issho/api/common"
)

// NewClaimsContext creates a new context from the given parent and adds the
// given JWT claims to it.
func NewClaimsContext(ctx context.Context, claims common.Claims) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

// ClaimsFromContext extracts JWT claims from the given context.
func ClaimsFromContext(ctx context.Context) (common.Claims, bool) {
	claims, ok := ctx.Value(claimsKey).(common.Claims)
	return claims, ok
}

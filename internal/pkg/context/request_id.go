package context

import (
	"context"

	"github.com/issho-ni/issho/api/common"
)

// NewRequestID generates a new request ID.
func NewRequestID() *common.UUID {
	return common.NewUUID()
}

// NewRequestIDContext creates a new context from the given parent and adds the given
// request ID to it.
func NewRequestIDContext(ctx context.Context, rid *common.UUID) context.Context {
	return context.WithValue(ctx, ridKey, rid)
}

// RequestIDFromContext extracts a request ID from the given context.
func RequestIDFromContext(ctx context.Context) (rid *common.UUID, ok bool) {
	*rid, ok = ctx.Value(ridKey).(common.UUID)
	return
}

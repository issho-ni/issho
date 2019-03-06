package requestid

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey int

const ridKey ctxKey = ctxKey(0)

// NewRequestID generates a new request ID.
func NewRequestID() uuid.UUID {
	return uuid.New()
}

// NewContext creates a new context from the given parent and adds the given
// request ID to it.
func NewContext(ctx context.Context, rid uuid.UUID) context.Context {
	return context.WithValue(ctx, ridKey, rid)
}

// FromContext extracts a request ID from the given context.
func FromContext(ctx context.Context) (uuid.UUID, bool) {
	rid, ok := ctx.Value(ridKey).(uuid.UUID)
	return rid, ok
}

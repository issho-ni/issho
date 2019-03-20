package context

import (
	"context"

	"github.com/google/uuid"
)

// NewRequestID generates a new request ID.
func NewRequestID() uuid.UUID {
	return uuid.New()
}

// NewRequestIDContext creates a new context from the given parent and adds the given
// request ID to it.
func NewRequestIDContext(ctx context.Context, rid uuid.UUID) context.Context {
	return context.WithValue(ctx, ridKey, rid)
}

// RequestIDFromContext extracts a request ID from the given context.
func RequestIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	rid, ok := ctx.Value(ridKey).(uuid.UUID)
	return rid, ok
}

package context

import (
	"context"

	"github.com/issho-ni/issho/internal/pkg/uuid"
)

// NewUserIDContext creates a new context from the given parent and adds the
// given user ID to it.
func NewUserIDContext(ctx context.Context, uid uuid.UUID) context.Context {
	return context.WithValue(ctx, uidKey, uid)
}

// UserIDFromContext extracts a user ID from the given context.
func UserIDFromContext(ctx context.Context) (uuid.UUID, bool) {
	uid, ok := ctx.Value(uidKey).(uuid.UUID)
	return uid, ok
}

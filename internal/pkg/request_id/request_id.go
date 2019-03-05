package requestid

import (
	"context"

	"github.com/google/uuid"
)

type ctxKey int

const ridKey ctxKey = ctxKey(0)

func NewRequestID() uuid.UUID {
	return uuid.New()
}

func NewContext(ctx context.Context, rid uuid.UUID) context.Context {
	return context.WithValue(ctx, ridKey, rid)
}

func FromContext(ctx context.Context) (uuid.UUID, bool) {
	rid, ok := ctx.Value(ridKey).(uuid.UUID)
	return rid, ok
}

package context

import (
	"context"
	"time"
)

// NewTimingContext creates a new context from the given parent and adds the
// specified time to it.
func NewTimingContext(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, timingKey, t)
}

// TimingFromContext extracts the timing from the given context.
func TimingFromContext(ctx context.Context) (t time.Time, ok bool) {
	t, ok = ctx.Value(timingKey).(time.Time)
	return
}

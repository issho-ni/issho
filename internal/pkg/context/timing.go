package context

import (
	"context"
	"time"
)

// NewTimingContext creates a new context from the given parent and adds the
// current time to it.
func NewTimingContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, timingKey, time.Now())
}

// TimingFromContext extracts the timing from the given context.
func TimingFromContext(ctx context.Context) (time.Time, bool) {
	t, ok := ctx.Value(timingKey).(time.Time)
	return t, ok
}

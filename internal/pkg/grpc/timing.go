package grpc

import (
	"context"
	"encoding/json"
	"time"

	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

const timingKey = "timing"
const startTimeKey = "start_time"

func appendTimingToOutgoingContext(ctx context.Context) context.Context {
	var err error
	var ok bool
	var t time.Time
	var value []byte

	if t, ok = icontext.TimingFromContext(ctx); !ok {
		return ctx
	}

	if value, err = json.Marshal(t); err != nil {
		return ctx
	}

	return metadata.AppendToOutgoingContext(ctx, timingKey, string(value))
}

func logTimingFromIncomingContext(ctx context.Context) context.Context {
	var md metadata.MD
	var ok bool
	var t time.Time
	var value []string

	if md, ok = metadata.FromIncomingContext(ctx); !ok {
		return ctx
	}

	if value = md.Get(timingKey); len(value) != 1 {
		return ctx
	}

	if err := json.Unmarshal([]byte(value[0]), &t); err != nil {
		return ctx
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		startTimeKey: t.Format(time.RFC3339Nano),
	})

	return icontext.NewTimingContext(ctx, t)
}

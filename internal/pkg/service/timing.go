package service

import (
	"context"
	"encoding/json"
	"time"

	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func appendTimingToOutgoingContext(ctx context.Context) context.Context {
	t, ok := icontext.TimingFromContext(ctx)
	if !ok {
		return ctx
	}

	value, err := json.Marshal(t)
	if err != nil {
		return ctx
	}

	return metadata.AppendToOutgoingContext(ctx, "timing", string(value))
}

func logTimingFromIncomingContext(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	value := md.Get("timing")
	if len(value) != 1 {
		return ctx
	}

	var t time.Time
	err := json.Unmarshal([]byte(value[0]), &t)
	if err != nil {
		return ctx
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		"start_time": t.Format(time.RFC3339Nano),
	})

	return icontext.NewTimingContext(ctx, t)
}

package graphql

import (
	"net/http"

	"github.com/issho-ni/issho/internal/pkg/context"
	"github.com/issho-ni/issho/internal/pkg/uuid"
)

const requestIDHeader = "X-Request-ID"

type requestIDHandler struct {
	http.Handler
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return &requestIDHandler{next}
}

func (h *requestIDHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rid, err := uuid.Parse(r.Header.Get(requestIDHeader))

	if err != nil {
		rid = uuid.New()
		r.Header.Set(requestIDHeader, rid.String())
	}

	ctx := context.NewRequestIDContext(r.Context(), rid)
	h.Handler.ServeHTTP(rw, r.WithContext(ctx))
}

package graphql

import (
	"net/http"
	"time"

	"github.com/issho-ni/issho/internal/pkg/context"
)

type timingHandler struct {
	http.Handler
}

func timingMiddleware(next http.Handler) http.Handler {
	return &timingHandler{next}
}

func (h *timingHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := context.NewTimingContext(r.Context(), time.Now())
	h.Handler.ServeHTTP(rw, r.WithContext(ctx))
}

package graphql

import (
	"net/http"

	"github.com/issho-ni/issho/api/common"
	"github.com/issho-ni/issho/internal/pkg/context"
)

const requestIDHeader = "X-Request-ID"

type requestIDHandler struct {
	http.Handler
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return &requestIDHandler{next}
}

func (h *requestIDHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rid, err := common.ParseUUID(r.Header.Get(requestIDHeader))
	if err != nil {
		rid = common.NewUUID()
		r.Header.Set(requestIDHeader, rid.String())
	}

	ctx := context.NewRequestIDContext(r.Context(), rid)
	h.Handler.ServeHTTP(rw, r.WithContext(ctx))
}

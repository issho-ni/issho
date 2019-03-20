package graphql

import (
	"math"
	"net/http"
	"time"

	"github.com/issho-ni/issho/internal/pkg/context"

	log "github.com/sirupsen/logrus"
)

type loggingHandler struct {
	http.Handler
}

func loggingMiddleware(next http.Handler) http.Handler {
	return &loggingHandler{next}
}

func (h *loggingHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	now := time.Now()

	w := &loggingResponseWriter{rw, 0, 0}
	h.Handler.ServeHTTP(w, r)

	since := math.Round(time.Since(now).Seconds()*1e6) / 1e3
	rid, _ := context.RequestIDFromContext(r.Context())

	log.WithFields(log.Fields{
		"http.method":      r.Method,
		"http.protocol":    r.Proto,
		"http.remote_addr": r.RemoteAddr,
		"http.service":     "graphql.Graphql",
		"http.size":        w.Size,
		"http.status":      w.StatusCode,
		"http.time_ms":     since,
		"request_id":       rid,
		"span.kind":        "server",
		"system":           "http",
	}).WithTime(now).Info(r.RequestURI)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	Size       int
	StatusCode int
}

func (rw *loggingResponseWriter) Write(b []byte) (int, error) {
	if rw.StatusCode == 0 {
		rw.WriteHeader(http.StatusOK)
	}

	size, err := rw.ResponseWriter.Write(b)
	rw.Size += size
	return size, err
}

func (rw *loggingResponseWriter) WriteHeader(statusCode int) {
	rw.ResponseWriter.WriteHeader(statusCode)
	rw.StatusCode = statusCode
}

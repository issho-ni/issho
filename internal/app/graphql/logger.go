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
	w := &loggingResponseWriter{rw, 0, 0}
	h.Handler.ServeHTTP(w, r)

	rid, _ := context.RequestIDFromContext(r.Context())

	entry := log.WithFields(log.Fields{
		"http.method":      r.Method,
		"http.protocol":    r.Proto,
		"http.remote_addr": r.RemoteAddr,
		"http.size":        w.Size,
		"http.status":      w.StatusCode,
		"request_id":       rid,
		"span.kind":        "server",
		"system":           "http",
	})

	if claims, ok := context.ClaimsFromContext(r.Context()); ok {
		entry = entry.WithField("user_id", claims.UserID)
	}

	if start, ok := context.TimingFromContext(r.Context()); ok {
		since := math.Round(time.Since(start).Seconds()*1e6) / 1e3
		entry = entry.WithField("http.time_ms", since).WithTime(start)
	}

	entry.Info(r.RequestURI)
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

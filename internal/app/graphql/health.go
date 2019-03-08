package graphql

import (
	"net/http"
)

func liveCheck(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

type readyChecker struct {
	*clientSet
}

func (s *readyChecker) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

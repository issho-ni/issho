package graphql

import (
	"net/http"
	"sync"

	"github.com/issho-ni/issho/internal/pkg/service"
)

func liveCheck(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

type readyChecker struct {
	healthCheckers []func() *service.GRPCStatus
}

func newReadyChecker(cs ClientSet) *readyChecker {
	healthCheckers := make([]func() *service.GRPCStatus, 0)

	for _, client := range cs.AllClients() {
		healthCheckers = append(healthCheckers, client.HealthCheck)
	}

	return &readyChecker{healthCheckers}
}

func (s *readyChecker) Length() int {
	return len(s.healthCheckers)
}

func (s *readyChecker) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	results := make(chan *service.GRPCStatus, len(s.healthCheckers))
	wg := sync.WaitGroup{}

	for _, checker := range s.healthCheckers {
		wg.Add(1)
		go func(c func() *service.GRPCStatus, wg *sync.WaitGroup) {
			defer wg.Done()
			results <- c()
		}(checker, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for status := range results {
		if !status.Result {
			rw.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
}

package graphql

import (
	"net/http"
	"sync"

	"github.com/issho-ni/issho/internal/pkg/grpc"
)

func liveCheck(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

type readyChecker struct {
	healthCheckers []func() *grpc.Status
}

func newReadyChecker(cs ClientSet) *readyChecker {
	healthCheckers := make([]func() *grpc.Status, 0)

	for _, client := range cs.AllClients() {
		healthCheckers = append(healthCheckers, client.HealthCheck)
	}

	return &readyChecker{healthCheckers}
}

func (s *readyChecker) Length() int {
	return len(s.healthCheckers)
}

func (s *readyChecker) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	results := make(chan *grpc.Status, len(s.healthCheckers))
	wg := sync.WaitGroup{}

	for _, checker := range s.healthCheckers {
		wg.Add(1)
		go func(c func() *grpc.Status, wg *sync.WaitGroup) {
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

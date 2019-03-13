package graphql

import (
	"net/http"
	"reflect"
	"sync"

	"github.com/issho-ni/issho/internal/pkg/service"
)

func liveCheck(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

type readyChecker struct {
	healthCheckers []func() bool
}

func newReadyChecker(cs *clientSet) *readyChecker {
	v := reflect.Indirect(reflect.ValueOf(cs))
	healthCheckers := make([]func() bool, 0)

	for i := 0; i < v.NumField(); i++ {
		client := reflect.Indirect(reflect.ValueOf(v.Field(i).Interface()))
		serviceClient := client.FieldByName("GRPCClient").Interface().(service.GRPCClient)
		healthCheckers = append(healthCheckers, serviceClient.HealthCheck)
	}

	return &readyChecker{healthCheckers}
}

func (s *readyChecker) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	done := make(chan bool, len(s.healthCheckers))
	wg := sync.WaitGroup{}

	for _, checker := range s.healthCheckers {
		wg.Add(1)
		go func(c func() bool, wg *sync.WaitGroup) {
			defer wg.Done()
			done <- c()
		}(checker, &wg)
	}

	go func() {
		wg.Wait()
		close(done)
	}()

	for result := range done {
		if !result {
			rw.WriteHeader(http.StatusServiceUnavailable)
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
}

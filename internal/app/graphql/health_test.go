package graphql

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/issho-ni/issho/api/kazoku"
	mock_kazoku "github.com/issho-ni/issho/api/mock/kazoku"
	mock_service "github.com/issho-ni/issho/internal/mock/pkg/service"
	"github.com/issho-ni/issho/internal/pkg/service"
)

func Test_liveCheck(t *testing.T) {
	rw := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/health", nil)

	type args struct {
		rw *httptest.ResponseRecorder
		r  *http.Request
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Sets 200 status", args{rw, r}, http.StatusOK},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			liveCheck(tt.args.rw, tt.args.r)
			if got := tt.args.rw.Result().StatusCode; got != http.StatusOK {
				t.Errorf("liveCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newReadyChecker(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockClient := &kazoku.Client{
		GRPCClient:   mock_service.NewMockGRPCClient(ctrl),
		KazokuClient: mock_kazoku.NewMockKazokuClient(ctrl),
	}
	mockClientSet := &clientSet{KazokuClient: mockClient}

	type args struct {
		cs *clientSet
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"New readyChecker", args{mockClientSet}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *newReadyChecker(tt.args.cs); got.Length() != tt.want {
				t.Errorf("newReadyChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readyChecker_ServeHTTP(t *testing.T) {
	ctrl := gomock.NewController(t)

	passingStatus := &service.GRPCStatus{Result: true, Error: nil}
	mockPassingClient := mock_service.NewMockGRPCClient(ctrl)
	mockPassingClient.EXPECT().HealthCheck().Return(passingStatus).AnyTimes()

	failingStatus := &service.GRPCStatus{Result: false, Error: fmt.Errorf("Error")}
	mockFailingClient := mock_service.NewMockGRPCClient(ctrl)
	mockFailingClient.EXPECT().HealthCheck().Return(failingStatus).AnyTimes()

	type healthCheckers []func() *service.GRPCStatus

	allPassing := healthCheckers{mockPassingClient.HealthCheck}
	allFailing := healthCheckers{mockFailingClient.HealthCheck}
	mixed := healthCheckers{mockPassingClient.HealthCheck, mockFailingClient.HealthCheck}

	tests := []struct {
		name           string
		healthCheckers healthCheckers
		want           int
	}{
		{"All passing", allPassing, http.StatusOK},
		{"All failing", allFailing, http.StatusServiceUnavailable},
		{"Mixed", mixed, http.StatusServiceUnavailable},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &readyChecker{
				healthCheckers: tt.healthCheckers,
			}

			rw := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/health", nil)

			s.ServeHTTP(rw, r)
			if got := rw.Result().StatusCode; got != tt.want {
				t.Errorf("ServeHTTP() = %v, want %v", got, tt.want)
			}
		})
	}
}
package service

import (
	"os"
	"reflect"
	"testing"
)

func Test_parseEnvironment(t *testing.T) {
	currEnv := os.Getenv("ISSHO_ENV")
	defer func() {
		os.Setenv("ISSHO_ENV", currEnv)
	}()

	tests := []struct {
		name string
		env  string
		want environment
	}{
		{"Development", "DEVELOPMENT", 0},
		{"Testing", "TESTING", 1},
		{"Production", "PRODUCTION", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("ISSHO_ENV", tt.env)

			if got := parseEnvironment(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseEnvironment() = %v, want %v", got, tt.want)
			}

			os.Unsetenv("ISSHO_ENV")
		})
	}
}

func Test_environment_Development(t *testing.T) {
	tests := []struct {
		name string
		e    environment
		want bool
	}{
		{"Development", 0, true},
		{"Testing", 1, false},
		{"Production", 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Development(); got != tt.want {
				t.Errorf("environment.Development() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_environment_Testing(t *testing.T) {
	tests := []struct {
		name string
		e    environment
		want bool
	}{
		{"Development", 0, false},
		{"Testing", 1, true},
		{"Production", 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Testing(); got != tt.want {
				t.Errorf("environment.Testing() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_environment_Production(t *testing.T) {
	tests := []struct {
		name string
		e    environment
		want bool
	}{
		{"Development", 0, false},
		{"Testing", 1, false},
		{"Production", 2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Production(); got != tt.want {
				t.Errorf("environment.Production() = %v, want %v", got, tt.want)
			}
		})
	}
}

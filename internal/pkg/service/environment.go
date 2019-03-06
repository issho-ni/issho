package service

import (
	"os"
	"strings"
)

type environment int

// Environment indicates the current application environment
var Environment environment

// Constants indicating the possible application environments
const (
	DEVELOPMENT environment = 0
	TESTING     environment = 1
	PRODUCTION  environment = 2
)

func init() {
	Environment = parseEnvironment()
}

func parseEnvironment() environment {
	switch env := os.Getenv("ISSHO_ENV"); strings.ToLower(env) {
	case "testing":
		return TESTING
	case "production":
		return PRODUCTION
	default:
		return DEVELOPMENT
	}
}

// Development indicates a development environment
func (e environment) Development() bool {
	return e == DEVELOPMENT
}

// Testing indicates a testing environment
func (e environment) Testing() bool {
	return e == TESTING
}

// Production indicates a production environment
func (e environment) Production() bool {
	return e == PRODUCTION
}

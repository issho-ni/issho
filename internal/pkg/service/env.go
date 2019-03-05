package service

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/credentials"
)

// Env defines the interface for the environment of a service's connection as a
// client to other services.
type Env interface {
	Creds() credentials.TransportCredentials
}

type env struct {
	creds credentials.TransportCredentials
}

func (e *env) Creds() credentials.TransportCredentials {
	return e.creds
}

// NewEnv generates a new service client environment.
func NewEnv(tlsCert string) Env {
	creds, err := credentials.NewClientTLSFromFile(tlsCert, "")

	if err != nil {
		log.Fatalf("Failed to generate credentials: %v", err)
	}

	return &env{creds}
}

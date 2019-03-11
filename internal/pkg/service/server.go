package service

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	defaultTLSCert = "localhost+2.pem"
	defaultTLSKey  = "localhost+2-key.pem"
)

// Server defines the interface of the server for a service.
type Server interface {
	StartServer()
	Serve(serve func() error)
}

// ServerConfig contains the configuration of the server for a service.
type ServerConfig struct {
	Name    string
	Port    string
	TLSCert string
	TLSKey  string
}

// NewServerConfig creates a new set of server configuration values.
func NewServerConfig(name string, defaultPort string) *ServerConfig {
	port := os.Getenv(fmt.Sprintf("%s_PORT", name))
	if port == "" {
		port = defaultPort
	}

	tlsCert := os.Getenv("TLS_CERT")
	if tlsCert == "" {
		tlsCert = defaultTLSCert
	}

	tlsKey := os.Getenv("TLS_KEY")
	if tlsKey == "" {
		tlsKey = defaultTLSKey
	}

	return &ServerConfig{name, port, tlsCert, tlsKey}
}

// Serve starts the server.
func (s *ServerConfig) Serve(serve func() error) {
	log.Infof("Starting service %s on port :%s", s.Name, s.Port)
	log.Fatal(serve())
}
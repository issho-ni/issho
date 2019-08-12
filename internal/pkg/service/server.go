package service

import (
	"fmt"
	"os"
	"strings"

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
	var port string
	var tlsCert string
	var tlsKey string

	if port = os.Getenv(strings.ToUpper(fmt.Sprintf("%s_port", name))); port == "" {
		port = defaultPort
	}

	if tlsCert = os.Getenv("TLS_CERT"); tlsCert == "" {
		tlsCert = defaultTLSCert
	}

	if tlsKey = os.Getenv("TLS_KEY"); tlsKey == "" {
		tlsKey = defaultTLSKey
	}

	setFormatter(name)
	return &ServerConfig{name, port, tlsCert, tlsKey}
}

// Serve starts the server.
func (s *ServerConfig) Serve(serve func() error) {
	log.WithField("port", s.Port).Info("Starting service")
	log.Fatal(serve())
}

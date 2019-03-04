package main

import (
	"os"

	"github.com/issho-ni/issho/internal/app/graphql"
	_ "github.com/issho-ni/issho/internal/pkg/issho"

	log "github.com/sirupsen/logrus"
)

const defaultPort = "8080"
const defaultTLSCert = "localhost+2.pem"
const defaultTLSKey = "localhost+2-key.pem"

func main() {
	port := os.Getenv("GRAPHQL_PORT")
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

	log.Printf("Starting graphql service")
	graphql.StartServer(port, tlsCert, tlsKey)
}

package main

import (
	"os"

	"github.com/issho-ni/issho/internal/app/ninshou"
	_ "github.com/issho-ni/issho/internal/pkg/issho"

	log "github.com/sirupsen/logrus"
)

const defaultPort = "8081"
const defaultTLSCert = "localhost+2.pem"
const defaultTLSKey = "localhost+2-key.pem"

func main() {
	port := os.Getenv("NINSHOU_PORT")
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

	log.Printf("Starting ninshou service")
	ninshou.StartServer(port, tlsCert, tlsKey)
}

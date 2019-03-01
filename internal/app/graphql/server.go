package graphql

import (
	"net/http"
	"os"

	"github.com/issho-ni/issho/internal/pkg/issho"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const defaultPort = "8080"
const defaultTLSCert = "localhost.pem"
const defaultTLSKey = "localhost-key.pem"

// StartServer starts the HTTP server for the GraphQL endpoint service.
func StartServer() {
	port := os.Getenv("PORT")
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

	r := mux.NewRouter()

	if issho.Environment.Development() {
		r.Handle("/", handler.Playground("GraphQL playground", "/query"))
		log.Printf("connect to https://localhost:%s/ for GraphQL playground", port)
	}

	r.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}})))

	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServeTLS(":"+port, tlsCert, tlsKey, r))
}

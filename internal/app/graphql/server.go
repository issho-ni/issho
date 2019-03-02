package graphql

import (
	"net/http"

	"github.com/issho-ni/issho/internal/pkg/issho"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// StartServer starts the HTTP server for the GraphQL endpoint service.
func StartServer(port string, tlsCert string, tlsKey string) {
	r := mux.NewRouter()

	if issho.Environment.Development() {
		r.Handle("/", handler.Playground("GraphQL playground", "/query"))
		log.Printf("connect to https://localhost:%s/ for GraphQL playground", port)
	}

	r.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}})))

	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServeTLS(":"+port, tlsCert, tlsKey, r))
}

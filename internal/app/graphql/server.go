package graphql

import (
	"net/http"

	"github.com/issho-ni/issho/api/graphql"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/issho"
	"github.com/issho-ni/issho/internal/pkg/service"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// StartServer starts the HTTP server for the GraphQL endpoint service.
func StartServer(port string, tlsCert string, tlsKey string) {
	r := mux.NewRouter()

	if issho.Environment.Development() {
		r.Handle("/", handler.Playground("GraphQL playground", "/query"))
		log.Printf("Connect to https://localhost:%s/ for GraphQL playground", port)
	}

	env := service.NewEnv(tlsCert)
	ninshou := ninshou.NewClient(env)
	resolver := &Resolver{ninshou}

	r.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver})))

	r.Use(requestIDMiddleware)
	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServeTLS(":"+port, tlsCert, tlsKey, r))
}

package graphql

import (
	"net/http"

	"github.com/issho-ni/issho/api/graphql"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/youji"
	"github.com/issho-ni/issho/internal/pkg/service"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
)

type graphQLServer struct {
	*mux.Router
	*service.ServerConfig
}

// NewGraphQLServer creates a new HTTP handler for the GraphQL service.
func NewGraphQLServer(config *service.ServerConfig) service.Server {
	r := mux.NewRouter()

	if service.Environment.Development() {
		r.Handle("/", handler.Playground("GraphQL playground", "/query"))
	}

	env := service.NewClientConfig(config.TLSCert)
	resolver := &Resolver{
		ninshou.NewClient(env),
		youji.NewClient(env),
	}

	r.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver})))
	r.HandleFunc("/live", liveCheck)
	r.HandleFunc("/ready", readyCheck)

	r.Use(requestIDMiddleware)
	r.Use(loggingMiddleware)

	return &graphQLServer{r, config}
}

func (s *graphQLServer) serve() error {
	return http.ListenAndServeTLS(":"+s.ServerConfig.Port, s.ServerConfig.TLSCert, s.ServerConfig.TLSKey, s.Router)
}

func (s *graphQLServer) StartServer() {
	s.Serve(s.serve)
}

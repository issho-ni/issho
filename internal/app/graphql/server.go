package graphql

import (
	"net/http"

	"github.com/issho-ni/issho/api/graphql"
	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/youji"
	"github.com/issho-ni/issho/internal/pkg/service"

	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
)

type graphQLServer struct {
	*mux.Router
	*service.ServerConfig
	*clientSet
}

type clientSet struct {
	NinkaClient   *ninka.Client
	NinshouClient *ninshou.Client
	YoujiClient   *youji.Client
}

// NewGraphQLServer creates a new HTTP handler for the GraphQL service.
func NewGraphQLServer(config *service.ServerConfig) service.Server {
	r := mux.NewRouter()

	if service.Environment.Development() {
		r.Handle("/", handler.Playground("GraphQL playground", "/query"))
	}

	env := service.NewGRPCClientConfig(config.TLSCert)
	clients := &clientSet{
		ninka.NewClient(env),
		ninshou.NewClient(env),
		youji.NewClient(env),
	}

	server := &graphQLServer{r, config, clients}

	c := graphql.Config{Resolvers: &Resolver{clients}}
	c.Directives.Protected = protectedFieldDirective

	r.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(c)))
	r.HandleFunc("/live", liveCheck)
	r.Handle("/ready", newReadyChecker(clients))

	r.Use(timingMiddleware)
	r.Use(requestIDMiddleware)
	r.Use(server.authenticationMiddleware)
	r.Use(loggingMiddleware)

	return server
}

func (s *graphQLServer) serve() error {
	return http.ListenAndServeTLS(":"+s.ServerConfig.Port, s.ServerConfig.TLSCert, s.ServerConfig.TLSKey, s.Router)
}

func (s *graphQLServer) StartServer() {
	s.Serve(s.serve)
}

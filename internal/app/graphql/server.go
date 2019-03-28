package graphql

import (
	"net/http"

	"github.com/issho-ni/issho/api/graphql"
	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/youji"
	"github.com/issho-ni/issho/internal/pkg/service"

	"github.com/99designs/gqlgen-contrib/gqlapollotracing"
	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	r.HandleFunc("/live", liveCheck)

	env := service.NewGRPCClientConfig(config.TLSCert)
	clients := &clientSet{
		ninka.NewClient(env),
		ninshou.NewClient(env),
		youji.NewClient(env),
	}

	r.Handle("/ready", newReadyChecker(clients))

	server := &graphQLServer{r, config, clients}

	c := graphql.Config{Resolvers: &Resolver{clients}}
	c.Directives.Protected = protectedFieldDirective

	corsOptions := cors.Options{
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "OPTIONS", "POST"},
	}

	var options []handler.Option
	if service.Environment.Development() {
		r.Handle("/", handler.Playground("GraphQL playground", "/query"))

		options = append(options, handler.RequestMiddleware(gqlapollotracing.RequestMiddleware()))
		options = append(options, handler.Tracer(gqlapollotracing.NewTracer()))

		corsOptions.AllowedOrigins = []string{"https://localhost:9000"}
		corsOptions.Debug = true
	} else {
		options = append(options, handler.IntrospectionEnabled(false))
	}

	r.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(c), options...))

	r.Use(timingMiddleware)
	r.Use(requestIDMiddleware)
	r.Use(server.authenticationMiddleware)
	r.Use(loggingMiddleware)
	r.Use(cors.New(corsOptions).Handler)

	return server
}

func (s *graphQLServer) serve() error {
	return http.ListenAndServeTLS(":"+s.ServerConfig.Port, s.ServerConfig.TLSCert, s.ServerConfig.TLSKey, s.Router)
}

func (s *graphQLServer) StartServer() {
	s.Serve(s.serve)
}

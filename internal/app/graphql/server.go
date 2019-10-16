package graphql

import (
	"net/http"

	"github.com/issho-ni/issho/api/graphql"
	"github.com/issho-ni/issho/api/kazoku"
	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/shinninjou"
	"github.com/issho-ni/issho/api/youji"
	"github.com/issho-ni/issho/internal/pkg/grpc"
	"github.com/issho-ni/issho/internal/pkg/service"

	"github.com/99designs/gqlgen-contrib/gqlapollotracing"
	"github.com/99designs/gqlgen/handler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Server defines the structure of a server for the GraphQL service.
type Server struct {
	*mux.Router
	*service.ServerConfig
	*clientSet
}

// NewServer creates a new HTTP handler for the GraphQL service.
func NewServer(config *service.ServerConfig) service.Server {
	var options []handler.Option
	var s *Server

	s.ServerConfig = config

	r := mux.NewRouter()
	s.Router = r

	env := grpc.NewClientConfig(config.TLSCert)
	clients := &clientSet{
		KazokuClient:     kazoku.NewClient(env),
		NinkaClient:      ninka.NewClient(env),
		NinshouClient:    ninshou.NewClient(env),
		ShinninjouClient: shinninjou.NewClient(env),
		YoujiClient:      youji.NewClient(env),
	}
	s.clientSet = clients

	r.HandleFunc("/live", liveCheck)
	r.Handle("/ready", newReadyChecker(clients))

	c := graphql.Config{Resolvers: &Resolver{clients}}
	c.Directives.Protected = protectedFieldDirective

	corsOptions := cors.Options{
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "OPTIONS", "POST"},
	}

	if service.Environment.Development() {
		r.Handle("/", handler.Playground("GraphQL playground", "/query"))

		options = append(options, handler.RequestMiddleware(gqlapollotracing.RequestMiddleware()))
		options = append(options, handler.Tracer(gqlapollotracing.NewTracer()))

		corsOptions.AllowedOrigins = []string{"https://localhost:8080/", "https://localhost:9000"}
	} else {
		options = append(options, handler.IntrospectionEnabled(false))
	}

	r.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(c), options...))

	r.Use(timingMiddleware)
	r.Use(requestIDMiddleware)
	r.Use(s.authenticationMiddleware)
	r.Use(loggingMiddleware)
	r.Use(cors.New(corsOptions).Handler)

	return s
}

// StartServer provides the callback function to start the server.
func (s *Server) StartServer() {
	s.Serve(s.serve)
}

func (s *Server) serve() error {
	return http.ListenAndServeTLS(":"+s.ServerConfig.Port, s.ServerConfig.TLSCert, s.ServerConfig.TLSKey, s.Router)
}

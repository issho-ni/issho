package main

import (
	"github.com/issho-ni/issho/internal/app/graphql"
	_ "github.com/issho-ni/issho/internal/pkg/issho"
	"github.com/issho-ni/issho/internal/pkg/service"
)

const defaultPort = "8080"

func main() {
	config := service.NewServerConfig("graphql", defaultPort)
	server := graphql.NewGraphQLServer(config)
	server.StartServer()
}

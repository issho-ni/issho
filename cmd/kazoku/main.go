package main

import (
	"github.com/issho-ni/issho/internal/app/kazoku"
	"github.com/issho-ni/issho/internal/pkg/service"
)

const defaultPort = "8085"

func main() {
	config := service.NewServerConfig("kazoku", defaultPort)
	server := kazoku.NewServer(config)
	server.StartServer()
}

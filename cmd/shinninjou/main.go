package main

import (
	"github.com/issho-ni/issho/internal/app/shinninjou"
	"github.com/issho-ni/issho/internal/pkg/service"
)

const defaultPort = "8083"

func main() {
	config := service.NewServerConfig("shinninjou", defaultPort)
	server := shinninjou.NewServer(config)
	server.StartServer()
}

package main

import (
	"github.com/issho-ni/issho/internal/app/youji"
	"github.com/issho-ni/issho/internal/pkg/service"
)

const defaultPort = "8083"

func main() {
	config := service.NewServerConfig("youji", defaultPort)
	server := youji.NewYoujiServer(config)
	server.StartServer()
}

package main

import (
	"github.com/issho-ni/issho/internal/app/ninshou"
	"github.com/issho-ni/issho/internal/pkg/service"
)

const defaultPort = "8081"

func main() {
	config := service.NewServerConfig("ninshou", defaultPort)
	server := ninshou.NewNinshouServer(config)
	server.StartServer()
}

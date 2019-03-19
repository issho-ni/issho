package main

import (
	"github.com/issho-ni/issho/internal/app/ninka"
	"github.com/issho-ni/issho/internal/pkg/service"
)

const defaultPort = "8084"

func main() {
	config := service.NewServerConfig("ninka", defaultPort)
	server := ninka.NewNinkaServer(config)
	server.StartServer()
}

package ninshou

import (
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/service"

	log "github.com/sirupsen/logrus"
)

type ninshouServer struct {
	ninshou.NinshouServer
}

// StartServer starts the gRPC server for the Ninshou service.
func StartServer(port string, tlsCert string, tlsKey string) {
	lis, grpcServer := service.NewServer(port, tlsCert, tlsKey)
	ninshou.RegisterNinshouServer(grpcServer, &ninshouServer{})
	log.Fatal(grpcServer.Serve(lis))
}

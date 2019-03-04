package ninshou

import (
	"net"

	"github.com/issho-ni/issho/api/ninshou"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ninshouServer struct {
	ninshou.NinshouServer
}

func newServer() *ninshouServer {
	return &ninshouServer{}
}

// StartServer starts the gRPC server for the Ninshou service.
func StartServer(port string, tlsCert string, tlsKey string) {
	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	creds, err := credentials.NewServerTLSFromFile(tlsCert, tlsKey)

	if err != nil {
		log.Fatalf("Failed to generate server credentials: %v", err)
	}

	opts = append(opts, grpc.Creds(creds))

	logger := log.StandardLogger()
	logrusEntry := log.NewEntry(logger)

	grpc_logrus.ReplaceGrpcLogger(logrusEntry)
	opts = append(opts, grpc_middleware.WithUnaryServerChain(grpc_logrus.UnaryServerInterceptor(logrusEntry)))
	opts = append(opts, grpc_middleware.WithStreamServerChain(grpc_logrus.StreamServerInterceptor(logrusEntry)))

	grpcServer := grpc.NewServer(opts...)
	ninshou.RegisterNinshouServer(grpcServer, newServer())
	log.Fatal(grpcServer.Serve(lis))
}

package ninshou

import (
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"
)

type ninshouServer struct {
	service.GRPCServer
	mongoClient service.MongoClient
	ninshou.NinshouServer
}

// NewNinshouServer returns a new gRPC server for the Ninshou service.
func NewNinshouServer(config *service.ServerConfig) service.Server {
	server := &ninshouServer{}
	server.GRPCServer = service.NewGRPCServer(config, server)
	server.mongoClient = service.NewMongoClient(config.Name)
	return server
}

func (s *ninshouServer) RegisterServer(srv *grpc.Server) {
	ninshou.RegisterNinshouServer(srv, s)
}

func (s *ninshouServer) StartServer() {
	cancel := s.mongoClient.Connect()
	defer cancel()

	s.createIndexes()
	s.GRPCServer.StartServer()
}

func (s *ninshouServer) createIndexes() {
	index := mongo.IndexModel{}
	index.Keys = bsonx.Doc{{Key: "email", Value: bsonx.Int32(1)}}
	index.Options = options.Index().SetUnique(true)

	s.mongoClient.CreateIndexes(service.NewIndexSet("users", index))
}

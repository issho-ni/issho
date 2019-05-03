package shinninjou

import (
	"github.com/issho-ni/issho/api/shinninjou"
	"github.com/issho-ni/issho/internal/pkg/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"
)

type shinninjouServer struct {
	service.GRPCServer
	mongoClient service.MongoClient
	shinninjou.ShinninjouServer
}

// NewShinninjouServer returns a new gRPC server for the Shinninjou service.
func NewShinninjouServer(config *service.ServerConfig) service.Server {
	server := &shinninjouServer{}
	server.GRPCServer = service.NewGRPCServer(config, server)
	server.mongoClient = service.NewMongoClient(config.Name)
	return server
}

func (s *shinninjouServer) RegisterServer(srv *grpc.Server) {
	shinninjou.RegisterShinninjouServer(srv, s)
}

func (s *shinninjouServer) StartServer() {
	cancel := s.mongoClient.Connect()
	defer cancel()

	s.createIndexes()
	s.GRPCServer.StartServer()
}

func (s *shinninjouServer) createIndexes() {
	index := mongo.IndexModel{}
	index.Keys = bsonx.Doc{{Key: "userid", Value: bsonx.Int32(1)}, {Key: "credentialtype", Value: bsonx.Int32(1)}}
	index.Options = options.Index().SetUnique(true)

	s.mongoClient.CreateIndexes(service.NewIndexSet("credentials", index))
}

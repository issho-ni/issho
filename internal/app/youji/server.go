package youji

import (
	"github.com/issho-ni/issho/api/youji"
	"github.com/issho-ni/issho/internal/pkg/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"
)

type youjiServer struct {
	service.GRPCServer
	mongoClient service.MongoClient
	youji.YoujiServer
}

// NewYoujiServer returns a new gRPC server for the Youji service.
func NewYoujiServer(config *service.ServerConfig) service.Server {
	server := &youjiServer{}
	server.GRPCServer = service.NewGRPCServer(config, server)
	server.mongoClient = service.NewMongoClient(config.Name)
	return server
}

func (s *youjiServer) RegisterServer(srv *grpc.Server) {
	youji.RegisterYoujiServer(srv, s)
}

func (s *youjiServer) StartServer() {
	cancel := s.mongoClient.Connect()
	defer cancel()

	s.createIndexes()
	s.GRPCServer.StartServer()
}

func (s *youjiServer) createIndexes() {
	index := mongo.IndexModel{}
	index.Keys = bsonx.Doc{{Key: "userid", Value: bsonx.Int32(1)}}

	s.mongoClient.CreateIndexes(service.NewIndexSet("todos", index))
}

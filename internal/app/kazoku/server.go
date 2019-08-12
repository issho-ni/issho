package kazoku

import (
	"github.com/issho-ni/issho/api/kazoku"
	"github.com/issho-ni/issho/internal/pkg/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"
)

type kazokuServer struct {
	service.GRPCServer
	mongoClient service.MongoClient
	kazoku.KazokuServer
}

// NewKazokuServer returns a new gRPC server for the Kazoku service.
func NewKazokuServer(config *service.ServerConfig) service.Server {
	server := &kazokuServer{}
	server.GRPCServer = service.NewGRPCServer(config, server)
	server.mongoClient = service.NewMongoClient(config.Name)
	return server
}

func (s *kazokuServer) RegisterServer(srv *grpc.Server) {
	kazoku.RegisterKazokuServer(srv, s)
}

func (s *kazokuServer) StartServer() {
	cancel := s.mongoClient.Connect()
	defer cancel()

	s.createIndexes()
	s.GRPCServer.StartServer()
}

func (s *kazokuServer) createIndexes() {
	userAccountsIndex := mongo.IndexModel{}
	userAccountsIndex.Keys = bsonx.Doc{
		{Key: "accountid", Value: bsonx.Int32(1)},
		{Key: "userid", Value: bsonx.Int32(1)},
	}
	userAccountsIndex.Options = options.Index().SetUnique(true)

	s.mongoClient.CreateIndexes(service.NewIndexSet("useraccounts", userAccountsIndex))
}

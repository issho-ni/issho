package kazoku

import (
	"github.com/issho-ni/issho/api/kazoku"
	"github.com/issho-ni/issho/internal/pkg/grpc"
	"github.com/issho-ni/issho/internal/pkg/mongo"
	"github.com/issho-ni/issho/internal/pkg/service"

	mmongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	ggrpc "google.golang.org/grpc"
)

// Server defines the structure of a server for the Kazoku service.
type Server struct {
	*grpc.Server
	kazoku.KazokuServer
}

// NewServer returns a new gRPC server for the Kazoku service.
func NewServer(config *service.ServerConfig) service.Server {
	var s *Server
	s.Server = grpc.NewServer(config, s)

	return s
}

// RegisterServer registers the gRPC server as a Kazoku service handler.
func (s *Server) RegisterServer(srv *ggrpc.Server) {
	kazoku.RegisterKazokuServer(srv, s)
}

// StartServer initializes the MongoDB connection and database and starts the server.
func (s *Server) StartServer() {
	s.defineIndexes()
	s.Server.StartServer()
}

func (s *Server) defineIndexes() {
	var userAccountsIndex mmongo.IndexModel
	userAccountsIndex.Keys = bsonx.Doc{
		{Key: "accountid", Value: bsonx.Int32(1)},
		{Key: "userid", Value: bsonx.Int32(1)},
	}
	userAccountsIndex.Options = options.Index().SetUnique(true)

	s.MongoClient.DefineIndexes(mongo.NewIndexSet("useraccounts", userAccountsIndex))
}

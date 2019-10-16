package ninshou

import (
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/grpc"
	"github.com/issho-ni/issho/internal/pkg/mongo"
	"github.com/issho-ni/issho/internal/pkg/service"

	mmongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	ggrpc "google.golang.org/grpc"
)

// Server defines the structure of a server for the Ninshou service.
type Server struct {
	*grpc.Server
	ninshou.NinshouServer
}

// NewServer returns a new gRPC server for the Ninshou service.
func NewServer(config *service.ServerConfig) *Server {
	var server *Server
	server.Server = grpc.NewServer(config, server)
	return server
}

// RegisterServer registers the gRPC server as a Ninshou service handler.
func (s *Server) RegisterServer(srv *ggrpc.Server) {
	ninshou.RegisterNinshouServer(srv, s)
}

// StartServer initializes the MongoDB connection and database and starts the server.
func (s *Server) StartServer() {
	s.defineIndexes()
	s.Server.StartServer()
}

func (s *Server) defineIndexes() {
	var index mmongo.IndexModel
	index.Keys = bsonx.Doc{{Key: "email", Value: bsonx.Int32(1)}}
	index.Options = options.Index().SetUnique(true)

	s.MongoClient.DefineIndexes(mongo.NewIndexSet("users", index))
}

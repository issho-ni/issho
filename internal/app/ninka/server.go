package ninka

import (
	"encoding/base64"
	"os"

	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/grpc"
	"github.com/issho-ni/issho/internal/pkg/mongo"
	"github.com/issho-ni/issho/internal/pkg/service"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/pascaldekloe/jwt"
	log "github.com/sirupsen/logrus"
	mmongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	ggrpc "google.golang.org/grpc"
)

// Server defines the structure of a server for the Kazoku service.
type Server struct {
	*grpc.Server
	ninka.NinkaServer
	ninshouClient *ninshou.Client
	secret        []byte
}

// Claims represents a set of claims with the token ID parsed into a UUID.
type Claims struct {
	*jwt.Claims
	ID     uuid.UUID
	UserID uuid.UUID
}

// NewServer returns a new gRPC server for the Ninka service.
func NewServer(config *service.ServerConfig) *Server {
	var server *Server
	server.Server = grpc.NewServer(config, server)

	env := grpc.NewClientConfig(config.TLSCert)
	server.ninshouClient = ninshou.NewClient(env)

	secret := os.Getenv("NINKA_JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT secret must be set")
	} else if decoded, err := base64.StdEncoding.DecodeString(secret); err != nil {
		log.WithField("err", err).Fatal("Could not decode JWT secret")
	} else {
		server.secret = decoded
	}

	return server
}

// RegisterServer registers the gRPC server as a Ninka service handler.
func (s *Server) RegisterServer(srv *ggrpc.Server) {
	ninka.RegisterNinkaServer(srv, s)
}

// StartServer initializes the MongoDB connection and database and starts the server.
func (s *Server) StartServer() {
	s.defineIndexes()
	s.Server.StartServer()
}

func (s *Server) defineIndexes() {
	var tokenIDIndex mmongo.IndexModel
	tokenIDIndex.Keys = bsonx.Doc{{Key: "tokenid", Value: bsonx.Int32(1)}}
	tokenIDIndex.Options = options.Index().SetUnique(true)

	var expiresAtIndex mmongo.IndexModel
	expiresAtIndex.Keys = bsonx.Doc{{Key: "expiresat", Value: bsonx.Int32(1)}}
	expiresAtIndex.Options = options.Index().SetExpireAfterSeconds(0)

	s.MongoClient.DefineIndexes(mongo.NewIndexSet("invalid_tokens", tokenIDIndex, expiresAtIndex))
}

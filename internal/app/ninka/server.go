package ninka

import (
	"encoding/base64"
	"os"

	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/grpc"
	"github.com/issho-ni/issho/internal/pkg/service"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/pascaldekloe/jwt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	ggrpc "google.golang.org/grpc"
)

// Server defines the structure of a server for the Kazoku service.
type Server struct {
	service.Server
	mongoClient   service.MongoClient
	ninshouClient *ninshou.Client
	ninka.NinkaServer
	secret []byte
}

// Claims represents a set of claims with the token ID parsed into a UUID.
type Claims struct {
	ID     uuid.UUID
	UserID uuid.UUID
	*jwt.Claims
}

// NewServer returns a new gRPC server for the Ninka service.
func NewServer(config *service.ServerConfig) service.Server {
	var s *Server
	s.Server = grpc.NewServer(config, s)
	s.mongoClient = service.NewMongoClient(config.Name)

	env := grpc.NewClientConfig(config.TLSCert)
	s.ninshouClient = ninshou.NewClient(env)

	secret := os.Getenv("NINKA_JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT secret must be set")
	} else if decoded, err := base64.StdEncoding.DecodeString(secret); err != nil {
		log.WithField("err", err).Fatal("Could not decode JWT secret")
	} else {
		s.secret = decoded
	}

	return s
}

// RegisterServer registers the gRPC server as a Ninka service handler.
func (s *Server) RegisterServer(srv *ggrpc.Server) {
	ninka.RegisterNinkaServer(srv, s)
}

// StartServer initializes the MongoDB connection and database and starts the server.
func (s *Server) StartServer() {
	cancel := s.mongoClient.Connect()
	defer cancel()

	s.createIndexes()
	s.Server.StartServer()
}

func (s *Server) createIndexes() {
	tokenIDIndex := mongo.IndexModel{}
	tokenIDIndex.Keys = bsonx.Doc{{Key: "tokenid", Value: bsonx.Int32(1)}}
	tokenIDIndex.Options = options.Index().SetUnique(true)

	expiresAtIndex := mongo.IndexModel{}
	expiresAtIndex.Keys = bsonx.Doc{{Key: "expiresat", Value: bsonx.Int32(1)}}
	expiresAtIndex.Options = options.Index().SetExpireAfterSeconds(0)

	s.mongoClient.CreateIndexes(service.NewIndexSet("invalid_tokens", tokenIDIndex, expiresAtIndex))
}

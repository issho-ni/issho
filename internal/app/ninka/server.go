package ninka

import (
	"encoding/base64"
	"os"

	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/service"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/pascaldekloe/jwt"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"
)

type ninkaServer struct {
	mongoClient   service.MongoClient
	ninshouClient *ninshou.Client
	service.GRPCServer
	ninka.NinkaServer
	secret []byte
}

// Claims represents a set of claims with the token ID parsed into a UUID.
type Claims struct {
	ID     uuid.UUID
	UserID uuid.UUID
	*jwt.Claims
}

// NewNinkaServer returns a new gRPC server for the Ninka service.
func NewNinkaServer(config *service.ServerConfig) service.Server {
	server := &ninkaServer{}
	server.GRPCServer = service.NewGRPCServer(config, server)
	server.mongoClient = service.NewMongoClient(config.Name)

	env := service.NewGRPCClientConfig(config.TLSCert)
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

func (s *ninkaServer) RegisterServer(srv *grpc.Server) {
	ninka.RegisterNinkaServer(srv, s)
}

func (s *ninkaServer) StartServer() {
	cancel := s.mongoClient.Connect()
	defer cancel()

	s.createIndexes()
	s.GRPCServer.StartServer()
}

func (s *ninkaServer) createIndexes() {
	tokenIDIndex := mongo.IndexModel{}
	tokenIDIndex.Keys = bsonx.Doc{{Key: "tokenid", Value: bsonx.Int32(1)}}
	tokenIDIndex.Options = options.Index().SetUnique(true)

	expiresAtIndex := mongo.IndexModel{}
	expiresAtIndex.Keys = bsonx.Doc{{Key: "expiresat", Value: bsonx.Int32(1)}}
	expiresAtIndex.Options = options.Index().SetExpireAfterSeconds(0)

	s.mongoClient.CreateIndexes(service.NewIndexSet("invalid_tokens", tokenIDIndex, expiresAtIndex))
}

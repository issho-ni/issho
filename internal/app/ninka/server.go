package ninka

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/issho-ni/issho/api/ninka"
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
	service.GRPCServer
	mongoClient service.MongoClient
	ninka.NinkaServer
	secret []byte
}

// Claims represents a set of claims with the token ID parsed into a UUID.
type Claims struct {
	ID uuid.UUID
	*jwt.Claims
}

// NewNinkaServer returns a new gRPC server for the Ninka service.
func NewNinkaServer(config *service.ServerConfig) service.Server {
	server := &ninkaServer{}
	server.GRPCServer = service.NewGRPCServer(config, server)
	server.mongoClient = service.NewMongoClient(config.Name)
	server.secret = []byte(os.Getenv("NINKA_JWT_SECRET"))

	return server
}

func (s *ninkaServer) RegisterServer(srv *grpc.Server) {
	ninka.RegisterNinkaServer(srv, s)
}

func (s *ninkaServer) StartServer() {
	cancel := s.mongoClient.Connect()
	defer cancel()

	s.GRPCServer.StartServer()
}

func (s *ninkaServer) createIndexes() {
	log.Debugf("Creating indexes")

	tokenIDIndex := mongo.IndexModel{}
	tokenIDIndex.Keys = bsonx.Doc{{Key: "tokenid", Value: bsonx.Int32(1)}}
	tokenIDIndex.Options = options.Index().SetUnique(true)

	expiresAtIndex := mongo.IndexModel{}
	expiresAtIndex.Keys = bsonx.Doc{{Key: "expiresat", Value: bsonx.Int32(1)}}
	expiresAtIndex.Options = options.Index().SetExpireAfterSeconds(0)

	createOptions := options.CreateIndexes().SetMaxTime(10 * time.Second)
	invalidTokens := s.mongoClient.Database().Collection("invalid_tokens").Indexes()

	for _, index := range []mongo.IndexModel{tokenIDIndex, expiresAtIndex} {
		result, err := invalidTokens.CreateOne(context.Background(), index, createOptions)
		if err != nil {
			log.Fatalf("Could not create index: %v", err)
		}

		log.Debugf("Created index %s", result)
	}
}

func (s *ninkaServer) extractClaims(token *ninka.Token) (*Claims, error) {
	t := []byte(token.Token)

	claims, err := jwt.HMACCheck(t, s.secret)
	if err != nil {
		return nil, err
	} else if ok := claims.Valid(time.Now()); !ok {
		return nil, fmt.Errorf("JWT has expired or contains invalid claims")
	}

	tokenID, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, err
	}

	return &Claims{tokenID, claims}, nil
}

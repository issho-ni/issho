package shinninjou

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/shinninjou"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *shinninjouServer) CreateCredential(ctx context.Context, in *shinninjou.Credential) (*shinninjou.CredentialResponse, error) {
	ins, err := bson.Marshal(in)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := s.mongoClient.Database().Collection("credentials")
	_, err = collection.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}

	return &shinninjou.CredentialResponse{Success: true}, nil
}

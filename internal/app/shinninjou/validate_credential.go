package shinninjou

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/shinninjou"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (s *shinninjouServer) ValidateCredential(ctx context.Context, in *shinninjou.Credential) (*shinninjou.CredentialResponse, error) {
	result := &bson.D{}
	filter := bson.D{{Key: "userid", Value: in.UserID}, {Key: "credentialtype", Value: in.CredentialType}}

	collection := s.mongoClient.Database().Collection("credentials")
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}

	m := result.Map()
	hash, ok := m["credential"].(primitive.Binary)
	if !ok {
		return nil, fmt.Errorf("Could not retrieve bytes of bcrypt hash")
	}

	switch in.CredentialType {
	case shinninjou.CredentialType_PASSWORD:
		err = bcrypt.CompareHashAndPassword(hash.Data, in.Credential)
		if err != nil {
			return nil, err
		}
	}

	return &shinninjou.CredentialResponse{Success: true}, nil
}

package shinninjou

import (
	"context"

	"github.com/issho-ni/issho/api/shinninjou"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (s *shinninjouServer) ValidateCredential(ctx context.Context, in *shinninjou.Credential) (*shinninjou.CredentialResponse, error) {
	result := &shinninjou.Credential{}
	filter := bson.D{{Key: "userid", Value: in.UserID}, {Key: "credentialtype", Value: in.CredentialType}}

	collection := s.mongoClient.Collection("credentials")
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}

	switch in.CredentialType {
	case shinninjou.CredentialType_PASSWORD:
		err = bcrypt.CompareHashAndPassword(result.Credential, in.Credential)
		if err != nil {
			return nil, err
		}
	}

	return &shinninjou.CredentialResponse{Success: true}, nil
}

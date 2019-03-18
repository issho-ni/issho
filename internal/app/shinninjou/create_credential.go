package shinninjou

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/shinninjou"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (s *shinninjouServer) CreateCredential(ctx context.Context, in *shinninjou.CredentialRequest) (*shinninjou.CredentialResponse, error) {
	credential := &shinninjou.Credential{UserID: in.UserID, CredentialType: in.CredentialType}

	switch in.CredentialType {
	case shinninjou.CredentialType_PASSWORD:
		password, err := bcrypt.GenerateFromPassword([]byte(in.Credential), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		credential.EncryptedCredential = password
	}

	ins, err := bson.Marshal(credential)
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

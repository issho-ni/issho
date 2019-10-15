package shinninjou

import (
	"context"

	"github.com/issho-ni/issho/api/shinninjou"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// CreateCredential creates an authentication credential for the current user.
// If the credential is a password, a bcrypt hash is generated and stored.
func (s *Server) CreateCredential(ctx context.Context, in *shinninjou.Credential) (*shinninjou.CredentialResponse, error) {
	var err error
	var ins []byte
	var password []byte

	switch in.CredentialType {
	case shinninjou.CredentialType_PASSWORD:
		if password, err = bcrypt.GenerateFromPassword(in.Credential, bcrypt.DefaultCost); err != nil {
			return nil, err
		}

		in.Credential = password
	}

	if ins, err = bson.Marshal(in); err != nil {
		return nil, err
	}

	collection := s.MongoClient.Collection("credentials")
	if _, err = collection.InsertOne(ctx, ins); err != nil {
		return nil, err
	}

	return &shinninjou.CredentialResponse{Success: true}, nil
}

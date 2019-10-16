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
	switch in.CredentialType {
	case shinninjou.CredentialType_PASSWORD:
		password, err := bcrypt.GenerateFromPassword(in.Credential, bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		in.Credential = password
	}

	ins, err := bson.Marshal(in)
	if err != nil {
		return nil, err
	}

	collection := s.MongoClient.Collection("credentials")
	_, err = collection.InsertOne(ctx, ins)

	return &shinninjou.CredentialResponse{Success: err == nil}, err
}

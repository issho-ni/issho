package shinninjou

import (
	"context"

	"github.com/issho-ni/issho/api/shinninjou"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// ValidateCredential validates the given credential against the stored record.
func (s *Server) ValidateCredential(ctx context.Context, in *shinninjou.Credential) (response *shinninjou.CredentialResponse, err error) {
	response = new(shinninjou.CredentialResponse)

	filter := bson.D{{Key: "userid", Value: in.UserID}, {Key: "credentialtype", Value: in.CredentialType}}
	collection := s.MongoClient.Collection("credentials")

	var result *shinninjou.Credential
	if err = collection.FindOne(ctx, filter).Decode(result); err != nil {
		response.Success = false
		return
	}

	switch in.CredentialType {
	case shinninjou.CredentialType_PASSWORD:
		err = bcrypt.CompareHashAndPassword(result.Credential, in.Credential)
	}

	response.Success = err == nil
	return
}

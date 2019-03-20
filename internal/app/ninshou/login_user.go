package ninshou

import (
	"context"

	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/shinninjou"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *ninshouServer) LoginUser(ctx context.Context, in *ninshou.LoginRequest) (*ninshou.User, error) {
	result := &ninshou.User{}
	filter := bson.D{{Key: "email", Value: in.Email}}

	collection := s.mongoClient.Database().Collection("users")
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}

	credential := &shinninjou.Credential{
		UserID:         result.Id,
		CredentialType: shinninjou.CredentialType_PASSWORD,
		Credential:     []byte(in.Password),
	}

	_, err = s.ShinninjouClient.ValidateCredential(ctx, credential)
	if err != nil {
		return nil, err
	}

	return result, nil
}

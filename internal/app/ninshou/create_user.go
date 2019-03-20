package ninshou

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/shinninjou"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *ninshouServer) CreateUser(ctx context.Context, in *ninshou.NewUser) (*ninshou.User, error) {
	id := uuid.New()
	now := time.Now()
	user := &ninshou.User{Id: &id, Name: in.Name, Email: in.Email, CreatedAt: &now}

	ins, err := bson.Marshal(user)
	if err != nil {
		return nil, err
	}

	collection := s.mongoClient.Database().Collection("users")
	_, err = collection.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}

	credential := &shinninjou.Credential{
		UserID:         &id,
		CredentialType: shinninjou.CredentialType_PASSWORD,
		Credential:     []byte(in.Password),
	}

	_, err = s.ShinninjouClient.CreateCredential(ctx, credential)
	if err != nil {
		return nil, err
	}

	return user, nil
}

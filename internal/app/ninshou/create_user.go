package ninshou

import (
	"context"
	"errors"
	"time"

	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/shinninjou"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *ninshouServer) CreateUser(ctx context.Context, in *ninshou.NewUser) (*ninshou.User, error) {
	if in.Name == "" {
		return nil, errors.New("name can't be empty")
	}

	if in.Email == "" {
		return nil, errors.New("email can't be empty")
	}

	if in.Password.Size() == 0 {
		return nil, errors.New("password can't be empty")
	}

	id := uuid.New()
	now := time.Now()
	user := &ninshou.User{Id: &id, Name: in.Name, Email: in.Email, CreatedAt: &now}

	ins, err := bson.Marshal(user)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := s.mongoClient.Database().Collection("users")
	_, err = collection.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}

	credential := &shinninjou.Credential{UserID: &id, CredentialType: shinninjou.CredentialType_PASSWORD, Credential: []byte(*in.Password)}
	_, err = s.ShinninjouClient.CreateCredential(ctx, credential)
	if err != nil {
		return nil, err
	}

	return user, nil
}

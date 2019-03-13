package ninshou

import (
	"context"
	"errors"
	"time"

	"github.com/issho-ni/issho/api/ninshou"
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

	if in.Password == "" {
		return nil, errors.New("password can't be empty")
	}

	id := uuid.New()
	user := &ninshou.User{Id: &id, Name: in.Name, Email: in.Email}

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

	return user, nil
}

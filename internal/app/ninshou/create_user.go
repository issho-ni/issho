package ninshou

import (
	"context"
	"errors"
	"time"

	"github.com/issho-ni/issho/api/ninshou"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
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

	id := [16]byte(uuid.New())
	user := &ninshou.User{XId: id[:], Name: in.Name, Email: in.Email}

	ins, err := bson.Marshal(user)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := s.mongoClient.Database().Collection("users")
	log.Debug(user)
	_, err = collection.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}

	return user, nil
}

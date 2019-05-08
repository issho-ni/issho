package ninshou

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *ninshouServer) CreateUser(ctx context.Context, in *ninshou.User) (*ninshou.User, error) {
	id := uuid.New()
	in.Id = &id
	now := time.Now()
	in.CreatedAt = &now

	ins, err := bson.Marshal(in)
	if err != nil {
		return nil, err
	}

	collection := s.mongoClient.Collection("users")
	_, err = collection.InsertOne(ctx, ins)
	if err != nil {
		return nil, err
	}

	return in, nil
}

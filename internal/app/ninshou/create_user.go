package ninshou

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *ninshouServer) CreateUser(ctx context.Context, in *ninshou.User) (*ninshou.User, error) {
	var err error
	var ins []byte

	id := uuid.New()
	in.Id = &id
	now := time.Now()
	in.CreatedAt = &now

	if ins, err = bson.Marshal(in); err != nil {
		return nil, err
	}

	collection := s.mongoClient.Collection("users")
	if _, err = collection.InsertOne(ctx, ins); err != nil {
		return nil, err
	}

	return in, nil
}

package ninshou

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/ninshou"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *ninshouServer) GetUser(ctx context.Context, in *ninshou.User) (*ninshou.User, error) {
	filter := bson.M{}

	if in.Id != nil {
		filter["_id"] = in.Id
	} else if in.Email != "" {
		filter["email"] = in.Email
	} else {
		return nil, fmt.Errorf("No filter parameters specified")
	}

	result := &ninshou.User{}
	collection := s.mongoClient.Collection("users")

	if err := collection.FindOne(ctx, filter).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

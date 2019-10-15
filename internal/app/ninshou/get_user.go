package ninshou

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/ninshou"

	"go.mongodb.org/mongo-driver/bson"
)

// GetUser finds a user record by user ID or email address.
func (s *Server) GetUser(ctx context.Context, in *ninshou.User) (*ninshou.User, error) {
	var filter bson.M

	if in.Id != nil {
		filter["_id"] = in.Id
	} else if in.Email != "" {
		filter["email"] = in.Email
	} else {
		return nil, fmt.Errorf("No filter parameters specified")
	}

	var result *ninshou.User
	collection := s.MongoClient.Collection("users")

	if err := collection.FindOne(ctx, filter).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

package ninshou

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/ninshou"

	"go.mongodb.org/mongo-driver/bson"
)

// GetUser finds a user record by user ID or email address.
func (s *Server) GetUser(ctx context.Context, in *ninshou.User) (result *ninshou.User, err error) {
	filter := bson.M{}

	if in.Id != nil {
		filter["_id"] = in.Id
	} else if in.Email != "" {
		filter["email"] = in.Email
	} else {
		return nil, fmt.Errorf("No filter parameters specified")
	}

	collection := s.MongoClient.Collection("users")
	err = collection.FindOne(ctx, filter).Decode(result)

	return
}

package kazoku

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/kazoku"

	"go.mongodb.org/mongo-driver/bson"
)

// GetAccount finds an account by account ID.
func (s *Server) GetAccount(ctx context.Context, in *kazoku.Account) (result *kazoku.Account, err error) {
	var filter bson.M

	if in.Id != nil {
		filter["_id"] = in.Id
	} else {
		return nil, fmt.Errorf("No filter parameters specified")
	}

	collection := s.MongoClient.Collection("accounts")
	err = collection.FindOne(ctx, filter).Decode(result)
	return
}

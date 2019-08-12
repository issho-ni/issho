package kazoku

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/kazoku"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *kazokuServer) GetAccount(ctx context.Context, in *kazoku.Account) (*kazoku.Account, error) {
	filter := bson.M{}

	if in.Id != nil {
		filter["_id"] = in.Id
	} else {
		return nil, fmt.Errorf("No filter parameters specified")
	}

	result := &kazoku.Account{}
	collection := s.mongoClient.Collection("accounts")

	if err := collection.FindOne(ctx, filter).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

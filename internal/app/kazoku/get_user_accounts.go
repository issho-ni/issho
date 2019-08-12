package kazoku

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/kazoku"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *kazokuServer) GetUserAccounts(ctx context.Context, in *kazoku.UserAccount) (*kazoku.UserAccounts, error) {
	filter := bson.M{}
	if in.AccountID != nil {
		filter["accountid"] = in.AccountID
	} else if in.UserID != nil {
		filter["userid"] = in.UserID
	} else {
		return nil, fmt.Errorf("No filter parameters specified")
	}

	results := []*kazoku.UserAccount{}
	collection := s.mongoClient.Collection("useraccounts")

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		result := &kazoku.UserAccount{}
		if err := cur.Decode(result); err == nil {
			results = append(results, result)
		}
	}

	return &kazoku.UserAccounts{UserAccounts: results}, nil
}

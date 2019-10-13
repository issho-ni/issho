package kazoku

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/kazoku"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetUserAccounts finds all associated users for a given account ID, or all
// associated accounts for a given user ID.
func (s *Server) GetUserAccounts(ctx context.Context, in *kazoku.UserAccount) (*kazoku.UserAccounts, error) {
	var cur *mongo.Cursor
	var err error

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

	if cur, err = collection.Find(ctx, filter); err != nil {
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

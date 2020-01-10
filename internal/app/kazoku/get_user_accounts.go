package kazoku

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/kazoku"

	"go.mongodb.org/mongo-driver/bson"
)

// GetUserAccounts finds all associated users for a given account ID, or all
// associated accounts for a given user ID.
func (s *Server) GetUserAccounts(ctx context.Context, in *kazoku.UserAccount) (results *kazoku.UserAccounts, err error) {
	filter := bson.M{}

	if in.AccountID != nil {
		filter["accountid"] = in.AccountID
	} else if in.UserID != nil {
		filter["userid"] = in.UserID
	} else {
		return nil, fmt.Errorf("No filter parameters specified")
	}

	collection := s.MongoClient.Collection("useraccounts")
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	results = new(kazoku.UserAccounts)
	for cur.Next(ctx) {
		var result *kazoku.UserAccount
		if err = cur.Decode(result); err == nil {
			results.UserAccounts = append(results.UserAccounts, result)
		}
	}

	return
}

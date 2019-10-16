package kazoku

import (
	"context"

	"github.com/issho-ni/issho/api/kazoku"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateUserAccount creates a new association between a user and an account.
func (s *Server) CreateUserAccount(ctx context.Context, in *kazoku.UserAccount) (*kazoku.UserAccount, error) {
	ins, err := bson.Marshal(in)
	if err != nil {
		return nil, err
	}

	collection := s.MongoClient.Collection("useraccounts")
	_, err = collection.InsertOne(ctx, ins)
	return in, err
}

package kazoku

import (
	"context"

	"github.com/issho-ni/issho/api/kazoku"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *kazokuServer) CreateUserAccount(ctx context.Context, in *kazoku.UserAccount) (*kazoku.UserAccount, error) {
	var err error
	var ins []byte

	if ins, err = bson.Marshal(in); err != nil {
		return nil, err
	}

	collection := s.mongoClient.Collection("useraccounts")
	if _, err = collection.InsertOne(ctx, ins); err != nil {
		return nil, err
	}

	return in, nil
}

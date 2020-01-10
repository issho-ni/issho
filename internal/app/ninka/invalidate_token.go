package ninka

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/common"
	"github.com/issho-ni/issho/api/ninka"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// InvalidToken stores the token IDs of manually invalidated JWTs until their
// normal expiration time.
type InvalidToken struct {
	TokenID   *common.UUID
	ExpiresAt *time.Time
}

// InvalidateToken saves the given token ID in the list of invalidated tokens.
// The record expires at the same time that the token itself would expire. If
// the token is already invalid, nothing is done.
func (s *Server) InvalidateToken(ctx context.Context, in *common.Claims) (*ninka.TokenResponse, error) {
	var ins []byte

	if invalid, err := s.isTokenInvalid(ctx, in.TokenID); err != nil {
		return &ninka.TokenResponse{Success: false}, err
	} else if !invalid {
		invalidToken := &InvalidToken{
			TokenID:   in.TokenID,
			ExpiresAt: in.ExpiresAt,
		}

		if ins, err = bson.Marshal(invalidToken); err != nil {
			return &ninka.TokenResponse{Success: false}, err
		}

		collection := s.MongoClient.Collection("invalid_tokens")
		if _, err = collection.InsertOne(ctx, ins); err != nil {
			return &ninka.TokenResponse{Success: false}, err
		}
	}

	return &ninka.TokenResponse{Success: true}, nil
}

func (s *Server) isTokenInvalid(ctx context.Context, tokenID *common.UUID) (bool, error) {
	var result *InvalidToken
	filter := bson.D{{Key: "tokenid", Value: tokenID}}

	collection := s.MongoClient.Collection("invalid_tokens")
	if err := collection.FindOne(ctx, filter).Decode(result); err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

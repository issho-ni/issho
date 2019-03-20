package ninka

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// InvalidToken stores the token IDs of manually invalidated JWTs until their
// normal expiration time.
type InvalidToken struct {
	TokenID   uuid.UUID
	ExpiresAt time.Time
}

func (s *ninkaServer) InvalidateToken(ctx context.Context, in *ninka.Token) (*ninka.TokenResponse, error) {
	claims, err := s.extractClaims(in)
	if err != nil {
		return &ninka.TokenResponse{Success: false}, err
	}

	invalid, err := s.isTokenInvalid(claims.ID)
	if err != nil {
		return &ninka.TokenResponse{Success: false}, err
	} else if !invalid {
		invalidToken := &InvalidToken{
			TokenID:   claims.ID,
			ExpiresAt: claims.Expires.Time(),
		}

		ins, err := bson.Marshal(invalidToken)
		if err != nil {
			return &ninka.TokenResponse{Success: false}, err
		}

		collection := s.mongoClient.Database().Collection("invalid_tokens")
		_, err = collection.InsertOne(ctx, ins)
		if err != nil {
			return &ninka.TokenResponse{Success: false}, err
		}
	}

	return &ninka.TokenResponse{Success: true}, nil
}

func (s *ninkaServer) isTokenInvalid(tokenID uuid.UUID) (bool, error) {
	result := &InvalidToken{}
	filter := bson.D{{Key: "tokenid", Value: tokenID}}

	collection := s.mongoClient.Database().Collection("invalid_tokens")
	err := collection.FindOne(context.Background(), filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

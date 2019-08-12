package kazoku

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/kazoku"
	"github.com/issho-ni/issho/internal/pkg/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *kazokuServer) CreateAccount(ctx context.Context, in *kazoku.Account) (*kazoku.Account, error) {
	var err error
	var ins []byte

	id := uuid.New()
	in.Id = &id
	// TODO Get this timestamp from the context request timing
	now := time.Now()
	in.CreatedAt = &now
	expires := now.Add(365 * 24 * time.Hour)
	in.ExpiresAt = &expires

	if ins, err = bson.Marshal(in); err != nil {
		return nil, err
	}

	collection := s.mongoClient.Collection("accounts")
	if _, err = collection.InsertOne(ctx, ins); err != nil {
		return nil, err
	}

	userAccount := &kazoku.UserAccount{
		AccountID:       &id,
		UserID:          in.CreatedByUserID,
		Role:            kazoku.UserAccount_OWNER,
		CreatedByUserID: in.CreatedByUserID,
		CreatedAt:       &now,
	}

	if _, err = s.CreateUserAccount(ctx, userAccount); err != nil {
		return nil, err
	}

	return in, nil
}

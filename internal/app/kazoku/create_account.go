package kazoku

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/kazoku"
	icontext "github.com/issho-ni/issho/internal/pkg/context"
	"github.com/issho-ni/issho/internal/pkg/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *kazokuServer) CreateAccount(ctx context.Context, in *kazoku.Account) (*kazoku.Account, error) {
	var err error
	var ins []byte
	var ok bool
	var t time.Time

	if t, ok = icontext.TimingFromContext(ctx); !ok {
		t = time.Now()
	}

	id := uuid.New()
	in.Id = &id
	in.CreatedAt = &t
	expires := t.Add(365 * 24 * time.Hour)
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
		CreatedAt:       &t,
	}

	if _, err = s.CreateUserAccount(ctx, userAccount); err != nil {
		return nil, err
	}

	return in, nil
}

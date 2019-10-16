package kazoku

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/kazoku"
	icontext "github.com/issho-ni/issho/internal/pkg/context"
	"github.com/issho-ni/issho/internal/pkg/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateAccount creates a new account owned by the creating user ID.
func (s *Server) CreateAccount(ctx context.Context, in *kazoku.Account) (*kazoku.Account, error) {
	if t, ok := icontext.TimingFromContext(ctx); ok {
		*in.CreatedAt = t
	} else {
		*in.CreatedAt = time.Now()
	}

	*in.Id = uuid.New()
	*in.ExpiresAt = in.CreatedAt.Add(365 * 24 * time.Hour)

	ins, err := bson.Marshal(in)
	if err != nil {
		return nil, err
	}

	collection := s.MongoClient.Collection("accounts")
	if _, err := collection.InsertOne(ctx, ins); err != nil {
		return nil, err
	}

	userAccount := &kazoku.UserAccount{
		AccountID:       in.Id,
		UserID:          in.CreatedByUserID,
		Role:            kazoku.UserAccount_OWNER,
		CreatedByUserID: in.CreatedByUserID,
		CreatedAt:       in.CreatedAt,
	}

	_, err = s.CreateUserAccount(ctx, userAccount)
	return in, err
}

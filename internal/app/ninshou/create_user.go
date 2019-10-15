package ninshou

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/ninshou"
	icontext "github.com/issho-ni/issho/internal/pkg/context"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateUser creates a new user record.
func (s *Server) CreateUser(ctx context.Context, in *ninshou.User) (*ninshou.User, error) {
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

	if ins, err = bson.Marshal(in); err != nil {
		return nil, err
	}

	collection := s.MongoClient.Collection("users")
	if _, err = collection.InsertOne(ctx, ins); err != nil {
		return nil, err
	}

	return in, nil
}

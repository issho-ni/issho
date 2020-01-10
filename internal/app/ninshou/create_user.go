package ninshou

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/common"
	"github.com/issho-ni/issho/api/ninshou"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"go.mongodb.org/mongo-driver/bson"
)

// CreateUser creates a new user record.
func (s *Server) CreateUser(ctx context.Context, in *ninshou.User) (*ninshou.User, error) {
	in.Id = common.NewUUID()

	if t, ok := icontext.TimingFromContext(ctx); ok {
		*in.CreatedAt = t
	} else {
		*in.CreatedAt = time.Now()
	}

	ins, err := bson.Marshal(in)
	if err != nil {
		return nil, err
	}

	collection := s.MongoClient.Collection("users")
	_, err = collection.InsertOne(ctx, ins)

	return in, err
}

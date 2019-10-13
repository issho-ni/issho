package youji

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/issho-ni/issho/internal/pkg/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateTodo creates a new todo record.
func (s *Server) CreateTodo(ctx context.Context, in *youji.NewTodo) (*youji.Todo, error) {
	var err error
	var ins []byte
	var ok bool
	var t time.Time

	if t, ok = icontext.TimingFromContext(ctx); !ok {
		t = time.Now()
	}

	claims, _ := icontext.ClaimsFromContext(ctx)
	id := uuid.New()
	todo := &youji.Todo{Id: &id, UserID: claims.UserID, Text: in.GetText(), CreatedAt: &t}

	if ins, err = bson.Marshal(todo); err != nil {
		return nil, err
	}

	collection := s.mongoClient.Collection("todos")
	if _, err = collection.InsertOne(ctx, ins); err != nil {
		return nil, err
	}

	return todo, nil
}

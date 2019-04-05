package youji

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/issho-ni/issho/internal/pkg/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *youjiServer) CreateTodo(ctx context.Context, in *youji.NewTodo) (*youji.Todo, error) {
	claims, _ := icontext.ClaimsFromContext(ctx)
	id := uuid.New()
	now := time.Now()
	todo := &youji.Todo{Id: &id, UserID: claims.UserID, Text: in.GetText(), CreatedAt: &now}

	ins, err := bson.Marshal(todo)
	if err != nil {
		return nil, err
	}

	collection := s.mongoClient.Database().Collection("todos")
	if _, err = collection.InsertOne(ctx, ins); err != nil {
		return nil, err
	}

	return todo, nil
}

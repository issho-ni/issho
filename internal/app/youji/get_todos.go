package youji

import (
	"context"

	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *youjiServer) GetTodos(ctx context.Context, in *youji.GetTodosParams) (*youji.Todos, error) {
	results := []*youji.Todo{}

	claims, _ := icontext.ClaimsFromContext(ctx)
	filter := bson.D{{Key: "userid", Value: claims.UserID}}

	collection := s.mongoClient.Database().Collection("todos")
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		result := &youji.Todo{}
		if err := cur.Decode(result); err == nil {
			results = append(results, result)
		}
	}

	return &youji.Todos{Todos: results}, nil
}

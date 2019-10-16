package youji

import (
	"context"

	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"go.mongodb.org/mongo-driver/bson"
)

// GetTodos finds all todo records for the current user by user ID.
func (s *Server) GetTodos(ctx context.Context, in *youji.GetTodosParams) (results *youji.Todos, err error) {
	claims, _ := icontext.ClaimsFromContext(ctx)
	filter := bson.D{{Key: "userid", Value: claims.UserID}}

	collection := s.MongoClient.Collection("todos")
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	results = new(youji.Todos)
	for cur.Next(ctx) {
		var result *youji.Todo
		if err = cur.Decode(result); err == nil {
			results.Todos = append(results.Todos, result)
		}
	}

	return
}

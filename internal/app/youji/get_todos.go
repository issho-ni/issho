package youji

import (
	"context"

	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetTodos finds all todo records for the current user by user ID.
func (s *Server) GetTodos(ctx context.Context, in *youji.GetTodosParams) (*youji.Todos, error) {
	var results []*youji.Todo
	var cur *mongo.Cursor
	var err error

	claims, _ := icontext.ClaimsFromContext(ctx)
	filter := bson.D{{Key: "userid", Value: claims.UserID}}

	collection := s.MongoClient.Collection("todos")
	if cur, err = collection.Find(ctx, filter); err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var result *youji.Todo
		if err := cur.Decode(result); err == nil {
			results = append(results, result)
		}
	}

	return &youji.Todos{Todos: results}, nil
}

package youji

import (
	"context"

	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateTodo updates a todo record.
func (s *Server) UpdateTodo(ctx context.Context, in *youji.UpdateTodoParams) (*youji.Todo, error) {
	var update primitive.M

	claims, _ := icontext.ClaimsFromContext(ctx)
	filter := bson.D{{Key: "_id", Value: in.Id}, {Key: "userid", Value: claims.UserID}}

	collection := s.mongoClient.Collection("todos")
	result := collection.FindOne(ctx, filter)

	currentTodo := &youji.Todo{}
	if err := result.Decode(currentTodo); err != nil {
		return nil, err
	}

	if update = currentTodo.UpdateOperatorsFromParams(in); len(update) == 0 {
		return currentTodo, nil
	}

	opts := &options.FindOneAndUpdateOptions{}
	opts.SetReturnDocument(options.After)

	todo := &youji.Todo{}
	if err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

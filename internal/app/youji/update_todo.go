package youji

import (
	"context"

	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// UpdateTodo updates a todo record.
func (s *Server) UpdateTodo(ctx context.Context, in *youji.UpdateTodoParams) (todo *youji.Todo, err error) {
	var currentTodo *youji.Todo

	claims, _ := icontext.ClaimsFromContext(ctx)
	filter := bson.D{{Key: "_id", Value: in.Id}, {Key: "userid", Value: claims.UserID}}

	collection := s.MongoClient.Collection("todos")
	result := collection.FindOne(ctx, filter)
	if err := result.Decode(currentTodo); err != nil {
		return nil, err
	}

	update := currentTodo.UpdateOperatorsFromParams(in)
	if len(update) == 0 {
		return currentTodo, nil
	}

	var opts *options.FindOneAndUpdateOptions
	opts.SetReturnDocument(options.After)

	err = collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(todo)
	return
}

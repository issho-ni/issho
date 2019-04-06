package youji

import (
	"context"

	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *youjiServer) UpdateTodo(ctx context.Context, in *youji.UpdateTodoParams) (*youji.Todo, error) {
	claims, _ := icontext.ClaimsFromContext(ctx)
	filter := bson.D{{Key: "_id", Value: in.Id}, {Key: "userid", Value: claims.UserID}}

	collection := s.mongoClient.Database().Collection("todos")
	result := collection.FindOne(ctx, filter)

	currentTodo := &youji.Todo{}
	if err := result.Decode(currentTodo); err != nil {
		return nil, err
	}

	update := currentTodo.UpdateOperatorsFromParams(in)
	if len(update) == 0 {
		return currentTodo, nil
	}

	opts := &options.FindOneAndUpdateOptions{}
	opts.SetReturnDocument(options.After)
	result = collection.FindOneAndUpdate(ctx, filter, update, opts)

	todo := &youji.Todo{}
	if err := result.Decode(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

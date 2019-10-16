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
func (s *Server) CreateTodo(ctx context.Context, in *youji.NewTodo) (todo *youji.Todo, err error) {
	todo = new(youji.Todo)
	*todo.Id = uuid.New()
	todo.Text = in.GetText()

	if t, ok := icontext.TimingFromContext(ctx); ok {
		*todo.CreatedAt = t
	} else {
		*todo.CreatedAt = time.Now()
	}

	claims, _ := icontext.ClaimsFromContext(ctx)
	todo.UserID = claims.UserID

	ins, err := bson.Marshal(todo)
	if err != nil {
		return nil, err
	}

	collection := s.MongoClient.Collection("todos")
	_, err = collection.InsertOne(ctx, ins)
	return
}

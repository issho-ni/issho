//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"

	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/youji"

	"github.com/google/uuid"
)

type todoResolver struct{ Resolver interface{} }

// NewTodoResolver returns a new todo resolver.
func NewTodoResolver(r interface{}) TodoResolver {
	return &todoResolver{r}
}

func (r *todoResolver) ID(ctx context.Context, obj *youji.Todo) (string, error) {
	return bytesToUUIDString(obj.Id)
}

type userResolver struct{ Resolver interface{} }

// NewUserResolver returns a new user resolver.
func NewUserResolver(r interface{}) UserResolver {
	return &userResolver{r}
}

func (r *userResolver) ID(ctx context.Context, obj *ninshou.User) (string, error) {
	return bytesToUUIDString(obj.Id)
}

func bytesToUUIDString(bytes []byte) (string, error) {
	id, err := uuid.FromBytes(bytes)
	if err != nil {
		return "", err
	}

	return id.String(), err
}

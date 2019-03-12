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

func (r *todoResolver) _id(ctx context.Context, obj *youji.Todo) (string, error) {
	return bytesToUUIDString(obj.XId)
}

type userResolver struct{ Resolver interface{} }

// NewUserResolver returns a new user resolver.
func NewUserResolver(r interface{}) UserResolver {
	return &userResolver{r}
}

func (r *userResolver) _id(ctx context.Context, obj *ninshou.User) (string, error) {
	return bytesToUUIDString(obj.XId)
}

func bytesToUUIDString(bytes []byte) (string, error) {
	id, err := uuid.FromBytes(bytes)
	if err != nil {
		return "", err
	}

	return id.String(), err
}

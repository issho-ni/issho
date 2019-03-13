//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"

	"github.com/issho-ni/issho/api/graphql"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/youji"
)

// Resolver is the base type for GraphQL operation resolvers.
type Resolver struct {
	*clientSet
}

// Mutation returns a new mutation resolver.
func (r *Resolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}

// Query returns a new query resolver.
func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

// CreateTodo creates a new Todo.
func (r *mutationResolver) CreateTodo(ctx context.Context, input youji.NewTodo) (*youji.Todo, error) {
	return r.YoujiClient.CreateTodo(ctx, &input)
}

// CreateUser creates a new User.
func (r *mutationResolver) CreateUser(ctx context.Context, input ninshou.NewUser) (*ninshou.User, error) {
	return r.NinshouClient.CreateUser(ctx, &input)
}

type queryResolver struct{ *Resolver }

// Todos returns all Todo items.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*youji.Todo, error) {
	todos, err := r.YoujiClient.GetTodos(ctx, &youji.GetTodosParams{})
	return todos.GetTodos(), err
}

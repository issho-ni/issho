//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"

	"github.com/issho-ni/issho/api/graphql"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver is the base type for GraphQL operation resolvers.
type Resolver struct{}

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
func (r *mutationResolver) CreateTodo(ctx context.Context, input graphql.NewTodo) (*graphql.Todo, error) {
	panic("not implemented")
}

// CreateUser creates a new User.
func (r *mutationResolver) CreateUser(ctx context.Context, input graphql.NewUser) (*graphql.User, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

// Todos returns all Todo items.
func (r *queryResolver) Todos(ctx context.Context) ([]graphql.Todo, error) {
	panic("not implemented")
}

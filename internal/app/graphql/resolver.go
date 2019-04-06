//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"

	"github.com/issho-ni/issho/api/graphql"
	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"
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
func (r *mutationResolver) CreateUser(ctx context.Context, input ninshou.NewUser) (*graphql.LoginResponse, error) {
	user, err := r.NinshouClient.CreateUser(ctx, &input)
	if err != nil {
		return nil, err
	}

	return r.getLoginResponse(ctx, user)
}

// LoginUser attempts to authenticate a user given an email address and
// password.
func (r *mutationResolver) LoginUser(ctx context.Context, input ninshou.LoginRequest) (*graphql.LoginResponse, error) {
	user, err := r.NinshouClient.LoginUser(ctx, &input)
	if err != nil {
		return nil, err
	}

	return r.getLoginResponse(ctx, user)
}

func (r *mutationResolver) LogoutUser(ctx context.Context, _ *bool) (bool, error) {
	claims, _ := icontext.ClaimsFromContext(ctx)

	response, err := r.NinkaClient.InvalidateToken(ctx, &claims)
	if err != nil {
		return false, err
	}

	return response.Success, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input youji.UpdateTodoParams) (*youji.Todo, error) {
	return r.YoujiClient.UpdateTodo(ctx, &input)
}

func (r *mutationResolver) getLoginResponse(ctx context.Context, user *ninshou.User) (*graphql.LoginResponse, error) {
	tokenRequest := &ninka.TokenRequest{UserID: user.Id}

	token, err := r.NinkaClient.CreateToken(ctx, tokenRequest)
	if err != nil {
		return nil, err
	}

	return &graphql.LoginResponse{Token: token.Token, User: *user}, nil
}

type queryResolver struct{ *Resolver }

// Todos returns all Todo items.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*youji.Todo, error) {
	todos, err := r.YoujiClient.GetTodos(ctx, &youji.GetTodosParams{})
	return todos.GetTodos(), err
}

//go:generate go run github.com/99designs/gqlgen

package graphql

import (
	"context"

	"github.com/issho-ni/issho/api/graphql"
	"github.com/issho-ni/issho/api/kazoku"
	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/api/ninshou"
	"github.com/issho-ni/issho/api/shinninjou"
	"github.com/issho-ni/issho/api/youji"
	icontext "github.com/issho-ni/issho/internal/pkg/context"
)

// Resolver is the base type for GraphQL operation resolvers.
type Resolver struct {
	*clientSet
}

// Account returns a new account resolver.
func (r *Resolver) Account() graphql.AccountResolver {
	return &accountResolver{r}
}

// Mutation returns a new mutation resolver.
func (r *Resolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}

// Query returns a new query resolver.
func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

// User returns a new user resolver.
func (r *Resolver) User() graphql.UserResolver {
	return &userResolver{r}
}

// UserAccount returns a new user account resolver.
func (r *Resolver) UserAccount() graphql.UserAccountResolver {
	return &userAccountResolver{r}
}

type accountResolver struct{ *Resolver }

func (r *accountResolver) UserAccounts(ctx context.Context, obj *kazoku.Account) ([]*kazoku.UserAccount, error) {
	userAccounts, err := r.KazokuClient.GetUserAccounts(ctx, &kazoku.UserAccount{AccountID: obj.Id})
	return userAccounts.GetUserAccounts(), err
}

// CreatedBy returns the user object for the createdByUserID field on an
// Account.
func (r *accountResolver) CreatedBy(ctx context.Context, obj *kazoku.Account) (*ninshou.User, error) {
	return r.NinshouClient.GetUser(ctx, &ninshou.User{Id: obj.CreatedByUserID})
}

// UpdatedBy returns the user object for the updatedByUserID field on an
// Account.
func (r *accountResolver) UpdatedBy(ctx context.Context, obj *kazoku.Account) (*ninshou.User, error) {
	return r.NinshouClient.GetUser(ctx, &ninshou.User{Id: obj.UpdatedByUserID})
}

type mutationResolver struct{ *Resolver }

// CreateAccount creates a new Account.
func (r *mutationResolver) CreateAccount(ctx context.Context, input graphql.NewAccount) (loginResponse *graphql.LoginResponse, err error) {
	if loginResponse, err = r.CreateUser(ctx, *input.User); err != nil {
		return nil, err
	}

	_, err = r.KazokuClient.CreateAccount(ctx, &kazoku.Account{Name: input.Name, CreatedByUserID: loginResponse.User.Id})

	return loginResponse, err
}

// CreateTodo creates a new Todo.
func (r *mutationResolver) CreateTodo(ctx context.Context, input youji.NewTodo) (*youji.Todo, error) {
	return r.YoujiClient.CreateTodo(ctx, &input)
}

// CreateUser creates a new User.
func (r *mutationResolver) CreateUser(ctx context.Context, input graphql.NewUser) (*graphql.LoginResponse, error) {
	user, err := r.NinshouClient.CreateUser(ctx, &ninshou.User{Name: input.Name, Email: input.Email})
	if err != nil {
		return nil, err
	}

	if _, err := r.ShinninjouClient.CreateCredential(ctx, &shinninjou.Credential{
		UserID:         user.Id,
		CredentialType: shinninjou.CredentialType_PASSWORD,
		Credential:     []byte(input.Password),
	}); err != nil {
		return nil, err
	}

	return r.getLoginResponse(ctx, user)
}

// LoginUser attempts to authenticate a user given an email address and
// password.
func (r *mutationResolver) LoginUser(ctx context.Context, input graphql.LoginRequest) (*graphql.LoginResponse, error) {
	user, err := r.NinshouClient.GetUser(ctx, &ninshou.User{Email: input.Email})
	if err != nil {
		return nil, err
	}

	if _, err := r.ShinninjouClient.ValidateCredential(ctx, &shinninjou.Credential{
		UserID:         user.Id,
		CredentialType: shinninjou.CredentialType_PASSWORD,
		Credential:     []byte(input.Password),
	}); err != nil {
		return nil, err
	}

	return r.getLoginResponse(ctx, user)
}

func (r *mutationResolver) LogoutUser(ctx context.Context, _ *bool) (bool, error) {
	claims, _ := icontext.ClaimsFromContext(ctx)
	response, err := r.NinkaClient.InvalidateToken(ctx, &claims)

	return response.Success, err
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input youji.UpdateTodoParams) (*youji.Todo, error) {
	return r.YoujiClient.UpdateTodo(ctx, &input)
}

func (r *mutationResolver) getLoginResponse(ctx context.Context, user *ninshou.User) (*graphql.LoginResponse, error) {
	tokenRequest := &ninka.TokenRequest{UserID: user.Id}
	token, err := r.NinkaClient.CreateToken(ctx, tokenRequest)

	return &graphql.LoginResponse{Token: token.Token, User: user}, err
}

type queryResolver struct{ *Resolver }

// Todos returns all Todo items.
func (r *queryResolver) GetTodos(ctx context.Context) ([]*youji.Todo, error) {
	todos, err := r.YoujiClient.GetTodos(ctx, &youji.GetTodosParams{})
	return todos.GetTodos(), err
}

type userResolver struct{ *Resolver }

// UserAccounts returns the user account objects whose userID field matches the
// ID field on a User.
func (r *userResolver) UserAccounts(ctx context.Context, obj *ninshou.User) ([]*kazoku.UserAccount, error) {
	userAccounts, err := r.KazokuClient.GetUserAccounts(ctx, &kazoku.UserAccount{UserID: obj.Id})
	return userAccounts.GetUserAccounts(), err
}

type userAccountResolver struct{ *Resolver }

// Account returns the account object for the accountID field on a userAccount
// object.
func (r *userAccountResolver) Account(ctx context.Context, obj *kazoku.UserAccount) (*kazoku.Account, error) {
	return r.KazokuClient.GetAccount(ctx, &kazoku.Account{Id: obj.AccountID})
}

// User returns the user object for the userID field on a userAccount object.
func (r *userAccountResolver) User(ctx context.Context, obj *kazoku.UserAccount) (*ninshou.User, error) {
	return r.NinshouClient.GetUser(ctx, &ninshou.User{Id: obj.UserID})
}

// CreatedBy returns the user object for the createdByUserID field on a
// userAccount object.
func (r *userAccountResolver) CreatedBy(ctx context.Context, obj *kazoku.UserAccount) (*ninshou.User, error) {
	return r.NinshouClient.GetUser(ctx, &ninshou.User{Id: obj.CreatedByUserID})
}

// UpdatedBy returns the user object for the updatedByUserID field on a
// userAccount object.
func (r *userAccountResolver) UpdatedBy(ctx context.Context, obj *kazoku.UserAccount) (*ninshou.User, error) {
	return r.NinshouClient.GetUser(ctx, &ninshou.User{Id: obj.UpdatedByUserID})
}

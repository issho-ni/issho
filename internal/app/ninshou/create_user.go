package ninshou

import (
	"context"
	"errors"

	"github.com/issho-ni/issho/api/ninshou"
)

func (s *ninshouServer) CreateUser(ctx context.Context, in *ninshou.NewUser) (*ninshou.User, error) {
	if in.Name == "" {
		return nil, errors.New("name can't be empty")
	}

	if in.Email == "" {
		return nil, errors.New("email can't be empty")
	}

	if in.Password == "" {
		return nil, errors.New("password can't be empty")
	}

	return &ninshou.User{Name: in.Name, Email: in.Email}, nil
}

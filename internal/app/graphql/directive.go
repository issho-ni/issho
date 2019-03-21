package graphql

import (
	"context"
	"fmt"

	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/99designs/gqlgen/graphql"
)

func protectedFieldDirective(ctx context.Context, _ interface{}, next graphql.Resolver, authRequired bool) (interface{}, error) {
	_, ok := icontext.UserIDFromContext(ctx)

	if authRequired && !ok {
		return nil, fmt.Errorf("Authentication required")
	} else if !authRequired && ok {
		return nil, fmt.Errorf("Can't do this while logged in")
	}

	return next(ctx)
}

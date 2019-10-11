package graphql

import (
	"context"
	"fmt"

	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/99designs/gqlgen/graphql"
)

const errAuthenticationRequired = "ERR_AUTHENTICATION_REQUIRED"
const errAuthenticationForbidden = "ERR_AUTHENTICATION_FORBIDDEN"

func protectedFieldDirective(ctx context.Context, _ interface{}, next graphql.Resolver, authRequired bool) (interface{}, error) {
	if _, ok := icontext.ClaimsFromContext(ctx); authRequired && !ok {
		return nil, fmt.Errorf(errAuthenticationRequired)
	} else if !authRequired && ok {
		return nil, fmt.Errorf(errAuthenticationForbidden)
	}

	return next(ctx)
}

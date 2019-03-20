package graphql

import (
	"net/http"
	"regexp"

	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/internal/pkg/context"
)

const authenticationHeader = "Authentication"

type authenticationHandler struct {
	*graphQLServer
	http.Handler
	bearerExpression *regexp.Regexp
}

func (s *graphQLServer) authenticationMiddleware(next http.Handler) http.Handler {
	bearerExpression, _ := regexp.Compile(`Bearer (\S+)`)
	return &authenticationHandler{s, next, bearerExpression}
}

func (h *authenticationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	bearer := r.Header.Get(authenticationHeader)
	matches := h.bearerExpression.FindStringSubmatch(bearer)

	if len(matches) == 2 {
		token := &ninka.Token{Token: matches[1]}

		response, err := h.graphQLServer.NinkaClient.ValidateToken(r.Context(), token)
		if err == nil {
			ctx := context.NewUserIDContext(r.Context(), *response.UserID)
			r = r.WithContext(ctx)
		}
	}

	h.Handler.ServeHTTP(rw, r)
}
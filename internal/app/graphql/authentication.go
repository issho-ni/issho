package graphql

import (
	"net/http"
	"regexp"

	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/internal/pkg/context"
)

const authorizationHeader = "Authorization"

type authenticationHandler struct {
	*Server
	http.Handler
	bearerExpression *regexp.Regexp
}

func (s *Server) authenticationMiddleware(next http.Handler) http.Handler {
	bearerExpression, _ := regexp.Compile(`Bearer (\S+)`)
	return &authenticationHandler{
		Server:           s,
		Handler:          next,
		bearerExpression: bearerExpression,
	}
}

func (h *authenticationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	bearer := r.Header.Get(authorizationHeader)

	if matches := h.bearerExpression.FindStringSubmatch(bearer); len(matches) == 2 {
		token := &ninka.Token{Token: matches[1]}

		if response, err := h.Server.NinkaClient.ValidateToken(r.Context(), token); err == nil {
			ctx := context.NewClaimsContext(r.Context(), *response.Claims)
			r = r.WithContext(ctx)
		}
	}

	h.Handler.ServeHTTP(rw, r)
}

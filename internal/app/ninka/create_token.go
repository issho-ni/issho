package ninka

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/ninka"
	icontext "github.com/issho-ni/issho/internal/pkg/context"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/pascaldekloe/jwt"
	log "github.com/sirupsen/logrus"
)

// CreateToken creates and signs a new JWT for an authenticated user.
func (s *Server) CreateToken(ctx context.Context, in *ninka.TokenRequest) (token *ninka.Token, err error) {
	token = new(ninka.Token)

	t, ok := icontext.TimingFromContext(ctx)
	if !ok {
		t = time.Now()
	}

	var claims *jwt.Claims
	claims.ID = uuid.New().String()
	claims.Expires = jwt.NewNumericTime(t.Add(30 * 24 * time.Hour))
	claims.NotBefore = jwt.NewNumericTime(t)
	claims.Subject = in.UserID.String()

	tkn, err := claims.HMACSign(jwt.HS256, s.secret)
	if err != nil {
		return nil, err
	}

	token.Token = string(tkn)

	ctxlogrus.AddFields(ctx, log.Fields{
		"user_id": claims.Subject,
	})

	return
}

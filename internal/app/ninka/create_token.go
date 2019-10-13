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
func (s *Server) CreateToken(ctx context.Context, in *ninka.TokenRequest) (*ninka.Token, error) {
	var err error
	var ok bool
	var t time.Time
	var token []byte

	if t, ok = icontext.TimingFromContext(ctx); !ok {
		t = time.Now()
	}

	notBefore := t
	expires := t.Add(30 * 24 * time.Hour)

	claims := &jwt.Claims{}
	claims.ID = uuid.New().String()
	claims.Expires = jwt.NewNumericTime(expires)
	claims.NotBefore = jwt.NewNumericTime(notBefore)
	claims.Subject = in.UserID.String()

	if token, err = claims.HMACSign(jwt.HS256, s.secret); err != nil {
		return nil, err
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		"user_id": claims.Subject,
	})

	return &ninka.Token{Token: string(token)}, nil
}

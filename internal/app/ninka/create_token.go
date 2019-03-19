package ninka

import (
	"context"
	"time"

	"github.com/issho-ni/issho/api/ninka"
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/pascaldekloe/jwt"
)

func (s *ninkaServer) CreateToken(ctx context.Context, in *ninka.TokenRequest) (*ninka.Token, error) {
	now := time.Now()
	notBefore := now.Add(-time.Second)
	expires := now.Add(30 * 24 * time.Hour)

	claims := &jwt.Claims{}
	claims.ID = uuid.New().String()
	claims.Expires = jwt.NewNumericTime(expires)
	claims.NotBefore = jwt.NewNumericTime(notBefore)
	claims.Subject = in.UserID.String()

	token, err := claims.HMACSign(jwt.HS256, s.secret)
	if err != nil {
		return nil, err
	}

	return &ninka.Token{Token: string(token)}, nil
}
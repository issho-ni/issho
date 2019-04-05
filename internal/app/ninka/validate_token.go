package ninka

import (
	"context"
	"fmt"
	"time"

	"github.com/issho-ni/issho/api/common"
	"github.com/issho-ni/issho/api/ninka"
	icontext "github.com/issho-ni/issho/internal/pkg/context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/pascaldekloe/jwt"
	log "github.com/sirupsen/logrus"
)

func (s *ninkaServer) ValidateToken(ctx context.Context, in *ninka.Token) (*ninka.TokenResponse, error) {
	t, ok := icontext.TimingFromContext(ctx)
	if !ok {
		t = time.Now()
	}

	claims, err := s.extractClaims(in, t)
	if err != nil {
		return &ninka.TokenResponse{Success: false}, err
	}

	ctxlogrus.AddFields(ctx, log.Fields{
		"user_id": claims.UserID.String(),
	})

	invalid, err := s.isTokenInvalid(claims.TokenID)
	if err != nil {
		return &ninka.TokenResponse{Success: false}, err
	} else if invalid {
		return &ninka.TokenResponse{Success: false}, fmt.Errorf("JWT has been invalidated")
	}

	return &ninka.TokenResponse{Claims: claims, Success: true}, nil
}

func (s *ninkaServer) extractClaims(token *ninka.Token, t time.Time) (*common.Claims, error) {
	tt := []byte(token.Token)

	claims, err := jwt.HMACCheck(tt, s.secret)
	if err != nil {
		return nil, err
	} else if ok := claims.Valid(t); !ok {
		return nil, fmt.Errorf("JWT has expired or contains invalid claims")
	}

	return common.ClaimsFromJWT(claims)
}

package ninka

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/ninka"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	log "github.com/sirupsen/logrus"
)

func (s *ninkaServer) ValidateToken(ctx context.Context, in *ninka.Token) (*ninka.TokenResponse, error) {
	claims, err := s.extractClaims(in)
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

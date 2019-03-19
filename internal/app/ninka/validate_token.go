package ninka

import (
	"context"
	"fmt"

	"github.com/issho-ni/issho/api/ninka"
)

func (s *ninkaServer) ValidateToken(ctx context.Context, in *ninka.Token) (*ninka.TokenResponse, error) {
	claims, err := s.extractClaims(in)
	if err != nil {
		return &ninka.TokenResponse{Success: false}, err
	}

	invalid, err := s.isTokenInvalid(claims.ID)
	if err != nil {
		return &ninka.TokenResponse{Success: false}, err
	} else if invalid {
		return &ninka.TokenResponse{Success: false}, fmt.Errorf("JWT has been invalidated")
	}

	return &ninka.TokenResponse{UserID: &claims.UserID, Success: true}, nil
}

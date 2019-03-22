//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. common/common.proto

package common

import (
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/pascaldekloe/jwt"
)

// ClaimsFromJWT parses and returns the necessary fields from a jwt.Claims
// object.
func ClaimsFromJWT(claims *jwt.Claims) (*Claims, error) {
	tokenID, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, err
	}

	expires := claims.Expires.Time()
	return &Claims{ExpiresAt: &expires, TokenID: &tokenID, UserID: &userID}, nil
}

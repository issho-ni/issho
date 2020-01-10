//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. common/common.proto

package common

import (
	"github.com/issho-ni/issho/internal/pkg/uuid"

	"github.com/pascaldekloe/jwt"
)

// ClaimsFromJWT parses and returns the necessary fields from a jwt.Claims
// object.
func ClaimsFromJWT(claims *jwt.Claims) (c *Claims, err error) {
	c = new(Claims)
	*c.ExpiresAt = claims.Expires.Time()

	if *c.TokenID, err = uuid.Parse(claims.ID); err != nil {
		return nil, err
	}

	if *c.UserID, err = uuid.Parse(claims.Subject); err != nil {
		return nil, err
	}

	return
}

package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	ID       int64
	UserType int32
	jwt.RegisteredClaims
}

package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

var signingKey []byte

func Init(secret string) {
	signingKey = []byte(secret)
}

// Validate проверяет токен и возвращает claims
func Validate(tokenStr string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	parsed, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil || !parsed.Valid {
		return nil, err
	}
	return claims, nil
}

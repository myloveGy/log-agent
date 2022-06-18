package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	Data interface{} `json:"data,omitempty"`
}

func GenerateToken(data interface{}, secret string, expire time.Duration) (string, error) {
	claims := Claims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return token, err
}

func ParseToken(data interface{}, token, secret string) (*Claims, error) {
	item := &Claims{Data: data}
	if _, err := jwt.ParseWithClaims(token, item, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	}); err != nil {
		return nil, err
	}

	return item, nil
}

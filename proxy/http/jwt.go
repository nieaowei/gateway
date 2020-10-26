package proxy_http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

const (
	JwtSignKey  = "gateway"
	JwtExpireAt = 60 * 60
)

func JwtDecode(toekenStr string) (*jwt.StandardClaims, error) {
	token, err := jwt.ParseWithClaims(toekenStr, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JwtSignKey), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return nil, errors.New("claims error")
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}
	return claims, err
}

func JwtEncode(claims jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JwtSignKey))
}

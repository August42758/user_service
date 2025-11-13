package auth

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = "gferfeferferfre"

type IJWTMaker interface {
	CreateToken(claims jwt.Claims, secretKey string) (string, error)
	ParseToken(secretKey, tokenString string) (*jwt.RegisteredClaims, error)
}

type JWTMaker struct {
	SecretKey string
}

func (j *JWTMaker) CreateToken(claims jwt.Claims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JWTMaker) ParseToken(secretKey, tokenString string) (*jwt.RegisteredClaims, error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		//проверяем на ошибку валидности токена(exp, sub и тд)
		if errors.Is(err, jwt.ErrTokenInvalidClaims) {
			return nil, ErrInvalidToken
		}
		//иначе что-то с серверной стороны
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

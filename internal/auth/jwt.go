package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = "gferfeferferfre"
var ExpirationTime = time.Now().Add(24 * time.Hour)

type CustomClaims struct {
	UserID string
	jwt.RegisteredClaims
}

func NewCustomClaims(userId string, experationTime time.Time) *CustomClaims {
	return &CustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(experationTime),
		},
	}
}

func CreateToken(experationTime time.Time, userId, secretKey string) (string, error) {
	customClaims := NewCustomClaims(userId, experationTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString, secretKey string) (*CustomClaims, error) {
	customClaims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, customClaims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	return customClaims, nil
}

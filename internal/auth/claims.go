package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var ExpirationTime = time.Now().Add(1 * time.Minute)

func NewRegisteredClaims(userId string, experationTime time.Time) (*jwt.RegisteredClaims, error) {
	jti, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}

	return &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(experationTime),
		Subject:   userId,
		ID:        jti.String(),
	}, nil
}

package auth

import "errors"

var (
	ErrInvalidToken = errors.New("Неверный JWT токен")
)

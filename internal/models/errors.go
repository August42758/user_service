package models

import "errors"

var (
	ErrDuplicatedEmail = errors.New("Такая почта уже зарегистрирована")
	ErrUserDoesntExist = errors.New("Такой пользователь не зарегистрирован")
)

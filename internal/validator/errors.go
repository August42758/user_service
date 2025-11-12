package validator

import "errors"

var (
	ErrWrongEmailFormat    = errors.New("Неверно формат почты")
	ErrWrongPasswordFormat = errors.New("Неверный формат пароля")
)

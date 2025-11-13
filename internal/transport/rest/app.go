package rest

import (
	"userservice/internal/auth"
	"userservice/internal/models"
	"userservice/internal/validator"
)

type UserApp struct {
	Validator validator.IValidator
	UserModel models.IUserModel
	JWTMaker  auth.IJWTMaker
}

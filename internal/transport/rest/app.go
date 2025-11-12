package rest

import (
	"userservice/internal/models"
	"userservice/internal/validator"
)

type UserApp struct {
	Validator validator.IValidator
	UserModel models.IUserModel
}

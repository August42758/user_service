package dto

type RequestRegisterUserDTO struct {
	Name     string "json:name"
	Email    string "json:email"
	Password string "json:password"
}

type RequestLoginUserDTO struct {
	Email    string "json:email"
	Password string "json:password"
}

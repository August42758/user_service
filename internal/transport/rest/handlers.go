package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"userservice/internal/auth"
	"userservice/internal/dto"
	"userservice/internal/models"
	"userservice/internal/validator"

	"golang.org/x/crypto/bcrypt"
)

func (a *UserApp) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var registerUserDTO dto.RequestRegisterUserDTO

	//проверяем совпадение типа данных из JSON c GO
	if err := json.NewDecoder(r.Body).Decode(&registerUserDTO); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//валидируем email
	if err := a.Validator.MatchEmail(registerUserDTO.Email, validator.EmailRegex); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//валидируем password
	if err := a.Validator.CountMinAmountCharsInPassword(registerUserDTO.Password, validator.MinPasswordLen); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerUserDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//добавляем пользователя в БД
	if err := a.UserModel.Insert(registerUserDTO.Name, registerUserDTO.Email, string(hashedPassword)); err != nil {
		if errors.Is(err, models.ErrDuplicatedEmail) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Вы зарегестрированы"))

}

func (a *UserApp) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var loginUserDTO dto.RequestLoginUserDTO

	if err := json.NewDecoder(r.Body).Decode(&loginUserDTO); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//валидируем email
	if err := a.Validator.MatchEmail(loginUserDTO.Email, validator.EmailRegex); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//валидируем password
	if err := a.Validator.CountMinAmountCharsInPassword(loginUserDTO.Password, validator.MinPasswordLen); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//извлекаем пользователя из БД
	id, hashedPassword, err := a.UserModel.Select(loginUserDTO.Email)
	if err != nil {
		if errors.Is(err, models.ErrUserDoesntExist) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	//сравниваем пароли
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(loginUserDTO.Password)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//генерируем токен
	token, err := auth.CreateToken(auth.ExpirationTime, strconv.Itoa(id), auth.SecretKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// http.SetCookie(w, &http.Cookie{
	// 	Name: "access_token",
	// 	Value: token,

	// })

	w.Header().Set("access_token", token)
	w.Write([]byte("Вы успешно зашли"))
}

func (a *UserApp) HandlePing(w http.ResponseWriter, r *http.Request) {

}

package main

import (
	"log"
	"net/http"

	"userservice/internal/database"
	"userservice/internal/models"
	"userservice/internal/transport/rest"
	"userservice/internal/validator"
)

func main() {
	db, err := database.ConnectDB("postgres://postgres:12345@localhost:5432/userservice?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	userModel := &models.UserModel{
		Db: db,
	}

	validator := &validator.Validator{}

	app := rest.UserApp{
		Validator: validator,
		UserModel: userModel,
	}

	if err := http.ListenAndServe(":8000", app.GetRoutes()); err != nil {
		log.Fatal(err)
	}
}

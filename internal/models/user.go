package models

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type IUserModel interface {
	Insert(name string, email string, password string) error
	Select(email string) (int, string, error)
}

type UserModel struct {
	Db *sql.DB
}

func (um *UserModel) Insert(name string, email string, password string) error {
	stmt := "INSERT INTO user_entity (name, email, password) VALUES ($1, $2, $3)"

	if _, err := um.Db.Exec(stmt, name, email, password); err != nil {
		var sqlError *pq.Error
		if errors.As(err, &sqlError) {
			if sqlError.Code == "23505" {
				return ErrDuplicatedEmail
			}
			return err
		}
	}

	return nil
}

func (um *UserModel) Select(email string) (int, string, error) {
	stmt := "SELECT id, password FROM user_entity WHERE email = $1"

	var id int
	var password string
	if err := um.Db.QueryRow(stmt, email).Scan(&id, &password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", ErrUserDoesntExist
		}
		return 0, "", err
	}

	return id, password, nil
}

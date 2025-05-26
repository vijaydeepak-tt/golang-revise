package models

import (
	"errors"

	"example.com/go_events_api/pkgs/db"
	"example.com/go_events_api/pkgs/utils"
)

type User struct {
	ID       int64
	Email    string `bindings:"required"`
	Password string `bindings:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	hashed, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashed)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err

}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	isPwdValid := utils.CheckPasswordHash(u.Password, hashedPassword)

	if !isPwdValid {
		return errors.New("Invalid credentials")
	}

	return nil
}

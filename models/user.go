package models

import (
	"errors"

	"github.com/intellect-sam/event/db"
	"github.com/intellect-sam/event/utils"
)

type User struct {
	ID       int64
	Email    string `binding: "required"`
	password string `binding: "required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err

}

func (u User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)

	if err != nil {
		return errors.New("Crendentials invalid")

	}

	passwordIsValid := utils.CheckPasswordHash(u.password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Crendentials invalid")
	}

	return nil
}

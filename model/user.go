package model

import (
	"errors"
	"golearn/first-api/db"
	"golearn/first-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	pwHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, pwHash)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	user.ID = id
	return err
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (user *User) Validate() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var dbPw string
	err := row.Scan(&user.ID, &dbPw)
	if err != nil {
		return errors.New("You must provide a 'password' value.")
	}

	pwValid := utils.ValidatePw(dbPw, user.Password)
	if !pwValid {
		return errors.New("The provided credentials were invalid.")
	}

	return nil
}

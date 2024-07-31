package models

import (
	"example.com/simple-api/db"
	"example.com/simple-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// var users []User = []User{}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id

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
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, err
}

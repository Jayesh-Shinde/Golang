package models

import (
	"example.org/rest-api/db"
	"example.org/rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) CheckCredentails() bool {
	query := `select id,password from users where email=?`
	result := db.DB.QueryRow(query, u.Email)
	var hashedPassword string
	err := result.Scan(&u.Id, &hashedPassword)
	if err != nil {
		return false
	}
	return utils.CheckPasswordWithHash(u.Password, hashedPassword)
}

func (u User) Save() (*User, error) {
	query := `
	INSERT INTO users(email,password) values (?,?)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	u.Id = id
	return &u, nil
}

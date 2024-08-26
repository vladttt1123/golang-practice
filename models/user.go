package models

import (
	"errors"
	"eventBooking/db"
	"eventBooking/utils"
	"fmt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query :=
		`INSERT INTO users(email, password) 
		 VALUES (?, ?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer statement.Close()

	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := statement.Exec(user.Email, hashedPass)

	if err != nil {
		fmt.Println(err)
		fmt.Println("HERE2")
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		fmt.Println("HERE3")
		return err
	}
	user.ID = id
	fmt.Println("User Inserted successfully", user.ID)
	return err
}

func (user *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		// email not found
		return errors.New("Credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		// password wrong
		return errors.New("Credentials invalid")
	}

	return nil

}

func GetAllUsers() ([]User, error) { // return type is event slice
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		fmt.Println("HERE5")
		return nil, err
	}
	defer rows.Close()
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			fmt.Println("HERE6")
			fmt.Println(err)
			return nil, err
		}
		fmt.Println(user)
		users = append(users, user)
	}

	return users, nil

}

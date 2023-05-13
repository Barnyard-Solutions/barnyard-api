package database

import (
	"barnyard/api/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertUser(email, passKey string) error {
	stmt, err := db.Prepare("INSERT INTO user (USER_MAIL, USER_PASS) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Failed to prepare insert user statement: ", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, passKey)
	if err != nil {
		fmt.Println("Failed to execute insert user statement: ", err)
		return err
	}

	return nil
}

func SelectUser(token string) (models.SafeUser, error) {
	var user models.SafeUser
	query := `Select u.USER_ID, u.USER_MAIL from user as u
	INNER JOIN barnyard.user_key as uk ON uk.USER_ID = u.USER_ID AND uk.USER_KEY = ?`

	err := db.QueryRow(query, token).Scan(&user.ID, &user.Email)
	if err != nil {
		fmt.Println("Failed to execute select user statement: ", err)
		return user, err
	}

	return user, nil
}

func GenerateUserToken(email, passKey string) (string, error) {
	query := "SELECT GenerateKey(?, ?)"
	var token string

	err := db.QueryRow(query, email, passKey).Scan(&token)
	if err != nil {
		fmt.Println("Failed to generate user token: ", err)
		return "", err
	}

	if token == "" {
		fmt.Println("Failed to generate token")
		return "", nil
	}

	return token, nil
}

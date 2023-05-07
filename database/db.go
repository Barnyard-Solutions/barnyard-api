package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(connectionString string) error {
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Database connected successfully")
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Database connection closed")
	}
}

func InsertUser(email, passKey string) error {
	stmt, err := db.Prepare("INSERT INTO user (USER_MAIL, USER_PASS) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, passKey)
	if err != nil {
		return err
	}

	fmt.Println("User inserted successfully")
	return nil
}

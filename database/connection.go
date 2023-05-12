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
		fmt.Println("Failed to open database connection: ", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping database: ", err)
		return err
	}

	fmt.Println("Database connected successfully")
	return nil
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			fmt.Println("Failed to close database connection: ", err)
		} else {
			fmt.Println("Database connection closed")
		}
	}
}

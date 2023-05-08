package database

import (
	"barnyard/api/models"
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

	return nil
}

func GenerateUserToken(email, passKey string) (string, error) {
	query := "SELECT GenerateKey(?, ?)"
	var token string

	err := db.QueryRow(query, email, passKey).Scan(&token)
	if err != nil {
		return "", err
	}

	if token == "" {
		fmt.Println("Failed to generate token")
		return "", nil
	}

	return token, nil
}

func GetFeeds(token string) []models.Feed {
	query := `SELECT FEED_ID, FEED_NAME, MAX(permission) AS permission
	FROM (
	  SELECT fv.FEED_ID, f.FEED_NAME, fv.USER_PERMISSION_LEVEL AS permission
	  FROM barnyard.feed_viewer AS fv
	  INNER JOIN barnyard.feed AS f ON fv.FEED_ID = f.FEED_ID 
	  INNER JOIN barnyard.user_key AS uk ON uk.USER_ID = fv.USER_ID AND uk.USER_KEY =?
	
	  UNION
	
	  SELECT f.FEED_ID, f.FEED_NAME, 8 AS permission
	  FROM barnyard.feed AS f
	  INNER JOIN barnyard.user_key AS uk ON f.OWNER_ID = uk.USER_ID AND uk.USER_KEY = ?
	) AS subquery
	GROUP BY FEED_ID, FEED_NAME
	`

	rows, err := db.Query(query, token, token)
	if err != nil {
		fmt.Println("Failed to execute query:", err)
		return nil
	}
	defer rows.Close()

	feeds := make([]models.Feed, 0)

	for rows.Next() {
		var feed models.Feed
		err := rows.Scan(&feed.ID, &feed.Name, &feed.Permission)
		if err != nil {
			fmt.Println("Failed to scan row:", err)
			continue
		}
		feeds = append(feeds, feed)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error occurred while iterating over rows:", err)
		return nil
	}

	return feeds
}

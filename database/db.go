package database

import (
	"barnyard/api/models"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

var (
	db  *sql.DB
	log *logrus.Logger
)

func InitDB(connectionString string) error {
	var err error
	log = logrus.New() // Initialize the log variable

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Errorf("Failed to open database connection: %s", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Errorf("Failed to ping database: %s", err)
		return err
	}

	log.Println("Database connected successfully")
	return nil
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Errorf("Failed to close database connection: %s", err)
		} else {
			log.Println("Database connection closed")
		}
	}
}

func InsertUser(email, passKey string) error {
	stmt, err := db.Prepare("INSERT INTO user (USER_MAIL, USER_PASS) VALUES (?, ?)")
	if err != nil {
		log.Errorf("Failed to prepare insert user statement: %s", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email, passKey)
	if err != nil {
		log.Errorf("Failed to execute insert user statement: %s", err)
		return err
	}

	return nil
}

func GenerateUserToken(email, passKey string) (string, error) {
	query := "SELECT GenerateKey(?, ?)"
	var token string

	err := db.QueryRow(query, email, passKey).Scan(&token)
	if err != nil {
		log.Errorf("Failed to generate user token: %s", err)
		return "", err
	}

	if token == "" {
		fmt.Println("Failed to generate token")
		return "", nil
	}

	return token, nil
}

func SelectFeeds(token string) ([]models.Feed, error) {
	query := `SELECT FEED_ID, FEED_NAME, MAX(permission) AS permission
	FROM (
	  SELECT fv.FEED_ID, f.FEED_NAME, fv.USER_PERMISSION_LEVEL AS permission
	  FROM barnyard.feed_member AS fv
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
		log.Errorf("Failed to execute query: %s", err)
		return nil, err
	}
	defer rows.Close()

	feeds := make([]models.Feed, 0)

	for rows.Next() {
		var feed models.Feed
		err := rows.Scan(&feed.ID, &feed.Name, &feed.Permission)
		if err != nil {
			log.Errorf("Failed to scan row: %s", err)
			continue
		}
		feeds = append(feeds, feed)
	}

	if err = rows.Err(); err != nil {
		log.Errorf("Error occurred while iterating over rows: %s", err)
		return nil, err
	}

	return feeds, nil
}

func InsertFeed(token, feedName string) (bool, error) {
	stmt, err := db.Prepare(`INSERT INTO feed (FEED_NAME, OWNER_ID) 
	SELECT ?, uk.USER_ID FROM user_key as uk where uk.USER_KEY = ? `)
	if err != nil {
		log.Errorf("Failed to prepare insert feed statement: %s", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(feedName, token)
	if err != nil {
		log.Errorf("Failed to execute insert feed statement: %s", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Failed to retrieve rows affected: %s", err)
		return false, err
	}

	if rowsAffected == 1 {
		return true, nil
	}

	return false, nil
}

func DeleteFeed(token string, feedID int) (bool, error) {
	stmt, err := db.Prepare(`
	DELETE f
	FROM feed AS f
	INNER JOIN user_key AS uk ON uk.USER_ID = f.OWNER_ID AND uk.USER_KEY = ?
	WHERE f.FEED_ID = ?`)
	if err != nil {
		log.Errorf("Failed to prepare delete feed statement: %s", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(token, feedID)
	if err != nil {
		log.Errorf("Failed to execute delete feed statement: %s", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Failed to retrieve rows affected: %s", err)
		return false, err
	}

	if rowsAffected == 1 {
		return true, nil
	}

	return false, nil
}

func InsertEvent(token, name1, name2, date string, feedID int) (bool, error) {
	stmt, err := db.Prepare(`INSERT INTO event (NAME_1, NAME_2, DATE, FEED_ID) 
	SELECT ?, ?, ?, p.FEED_ID
		FROM (
			SELECT fv.FEED_ID
			FROM barnyard.feed_member AS fv
			INNER JOIN barnyard.feed AS f ON fv.FEED_ID = f.FEED_ID
			INNER JOIN barnyard.user_key AS uk ON uk.USER_ID = fv.USER_ID AND uk.USER_KEY = ?
			
			UNION
			
			SELECT f.FEED_ID
			FROM barnyard.feed AS f
			INNER JOIN barnyard.user_key AS uk ON f.OWNER_ID = uk.USER_ID AND uk.USER_KEY = ?
		) AS p
		WHERE p.FEED_ID = ?;
	
	`)
	if err != nil {
		log.Errorf("Failed to prepare insert event statement: %s", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name1, name2, date, token, token, feedID)
	if err != nil {
		log.Errorf("Failed to execute insert event statement: %s", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Failed to retrieve rows affected: %s", err)
		return false, err
	}

	if rowsAffected == 1 {
		return true, nil
	}

	return false, nil
}

func SelectEvents(token string, feedID int) ([]models.Event, error) {

	query := `
	SELECT e.EVENT_ID, e.NAME_1, e.NAME_2, e.DATE, p.FEED_ID
		FROM (
			SELECT fv.FEED_ID
			FROM barnyard.feed_member AS fv
			INNER JOIN barnyard.feed AS f ON fv.FEED_ID = f.FEED_ID
			INNER JOIN barnyard.user_key AS uk ON uk.USER_ID = fv.USER_ID AND uk.USER_KEY = ?
			
			UNION
			
			SELECT f.FEED_ID
			FROM barnyard.feed AS f
			INNER JOIN barnyard.user_key AS uk ON f.OWNER_ID = uk.USER_ID AND uk.USER_KEY = ?
		) AS p
		INNER JOIN event as e ON e.FEED_ID = p.FEED_ID
		WHERE p.FEED_ID = ?;
	`

	rows, err := db.Query(query, token, token, feedID)
	if err != nil {
		log.Errorf("Failed to execute query: %s", err)
		return nil, err
	}
	defer rows.Close()

	events := make([]models.Event, 0)

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Name1, &event.Name2, &event.Date, &event.FeedID)
		if err != nil {
			log.Errorf("Failed to scan row: %s", err)
			continue
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		log.Errorf("Error occurred while iterating over rows: %s", err)
		return nil, err
	}

	return events, nil

}

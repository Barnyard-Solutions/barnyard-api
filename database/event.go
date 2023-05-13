package database

import (
	"barnyard/api/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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
		fmt.Println("Failed to prepare insert event statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name1, name2, date, token, token, feedID)
	if err != nil {
		fmt.Println("Failed to execute insert event statement: ", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Failed to retrieve rows affected: ", err)
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
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	}
	defer rows.Close()

	events := make([]models.Event, 0)

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Name1, &event.Name2, &event.Date, &event.FeedID)
		if err != nil {
			fmt.Println("Failed to scan row: ", err)
			continue
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error occurred while iterating over rows: ", err)
		return nil, err
	}

	return events, nil

}

func DeleteEvent(token string, feedID, eventID int) (bool, error) {
	stmt, err := db.Prepare(`
	DELETE e
	FROM event AS e
	INNER JOIN feed AS f ON f.FEED_ID = e.FEED_ID
	INNER JOIN user_key AS uk ON uk.USER_ID = f.OWNER_ID AND uk.USER_KEY = ?
	WHERE f.FEED_ID = ? AND e.EVENT_ID = ?`)
	if err != nil {
		fmt.Println("Failed to prepare delete feed statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(token, feedID, eventID)
	if err != nil {
		fmt.Println("Failed to execute delete event statement: ", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Failed to retrieve rows affected: ", err)
		return false, err
	}

	if rowsAffected == 1 {
		return true, nil
	}

	return false, ErrNotFound
}

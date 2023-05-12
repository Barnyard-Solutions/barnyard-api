package database

import (
	"barnyard/api/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	}
	defer rows.Close()

	feeds := make([]models.Feed, 0)

	for rows.Next() {
		var feed models.Feed
		err := rows.Scan(&feed.ID, &feed.Name, &feed.Permission)
		if err != nil {
			fmt.Println("Failed to scan row: ", err)
			continue
		}
		feeds = append(feeds, feed)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error occurred while iterating over rows: ", err)
		return nil, err
	}

	return feeds, nil
}

func InsertFeed(token, feedName string) (bool, error) {
	stmt, err := db.Prepare(`INSERT INTO feed (FEED_NAME, OWNER_ID) 
	SELECT ?, uk.USER_ID FROM user_key as uk where uk.USER_KEY = ? `)
	if err != nil {
		fmt.Println("Failed to prepare insert feed statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(feedName, token)
	if err != nil {
		fmt.Println("Failed to execute insert feed statement: ", err)
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

func DeleteFeed(token string, feedID int) (bool, error) {
	stmt, err := db.Prepare(`
	DELETE f
	FROM feed AS f
	INNER JOIN user_key AS uk ON uk.USER_ID = f.OWNER_ID AND uk.USER_KEY = ?
	WHERE f.FEED_ID = ?`)
	if err != nil {
		fmt.Println("Failed to prepare delete feed statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(token, feedID)
	if err != nil {
		fmt.Println("Failed to execute delete feed statement: ", err)
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

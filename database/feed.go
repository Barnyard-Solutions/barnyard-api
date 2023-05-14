package database

import (
	"barnyard/api/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SelectFeeds(token string) ([]models.Feed, error) {
	query := `SELECT fmp.FEED_ID, f.FEED_NAME, fmp.permission from feed_member_permission AS fmp
	INNER JOIN user_key as uk ON uk.USER_ID = fmp.USER_ID AND uk.USER_KEY = ?
	INNER JOIN feed as f ON f.FEED_ID = fmp.FEED_ID
	`

	rows, err := db.Query(query, token)
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

func SelectFeedsWithSub(token, subscription string) ([]models.FeedSub, error) {
	query := `SELECT fmp.FEED_ID, f.FEED_NAME, fmp.permission, s.END_POINT IS NOT NULL from feed_member_permission AS fmp
	INNER JOIN user_key as uk ON uk.USER_ID = fmp.USER_ID AND uk.USER_KEY = ?
	INNER JOIN feed as f ON f.FEED_ID = fmp.FEED_ID
	LEFT JOIN subscription as s ON s.FEED_ID = fmp.FEED_ID AND s.USER_ID = fmp.USER_ID AND s.END_POINT = ?
	`

	rows, err := db.Query(query, token, subscription)
	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	}
	defer rows.Close()

	feeds := make([]models.FeedSub, 0)

	for rows.Next() {
		var feed models.FeedSub
		err := rows.Scan(&feed.ID, &feed.Name, &feed.Permission, &feed.Subscribed)
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

	return false, ErrNotFound
}

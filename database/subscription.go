package database

import (
	"barnyard/api/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertSubscription(token, subscription string, feedID int) (bool, error) {
	stmt, err := db.Prepare(`INSERT INTO subscription (USER_ID, FEED_ID, END_POINT) 
	SELECT p.USER_ID, p.FEED_ID, ?	
		FROM (
			SELECT fv.FEED_ID, uk.USER_ID
			FROM barnyard.feed_member AS fv
			INNER JOIN barnyard.feed AS f ON fv.FEED_ID = f.FEED_ID
			INNER JOIN barnyard.user_key AS uk ON uk.USER_ID = fv.USER_ID AND uk.USER_KEY = ?
			
			UNION
			
			SELECT f.FEED_ID, uk.USER_ID
			FROM barnyard.feed AS f
			INNER JOIN barnyard.user_key AS uk ON f.OWNER_ID = uk.USER_ID AND uk.USER_KEY = ?
		) AS p
		WHERE p.FEED_ID = ?;
	
	`)
	if err != nil {
		fmt.Println("Failed to prepare insert subscription statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(subscription, token, token, feedID)
	if err != nil {
		fmt.Println("Failed to execute insert subscription statement: ", err)
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

func SelectSubscription(token, subscription string, feedID int) ([]models.Subscription, error) {

	query := `
	SELECT p.USER_ID, p.FEED_ID, s.END_POINT
		FROM (
			SELECT fv.FEED_ID, uk.USER_ID
			FROM barnyard.feed_member AS fv
			INNER JOIN barnyard.feed AS f ON fv.FEED_ID = f.FEED_ID
			INNER JOIN barnyard.user_key AS uk ON uk.USER_ID = fv.USER_ID AND uk.USER_KEY = ?

			UNION

			SELECT f.FEED_ID, uk.USER_ID
			FROM barnyard.feed AS f
			INNER JOIN barnyard.user_key AS uk ON f.OWNER_ID = uk.USER_ID AND uk.USER_KEY = ?
		) AS p
		INNER JOIN subscription as s ON s.FEED_ID = p.FEED_ID AND s.USER_ID = p.USER_ID
		WHERE p.FEED_ID = ? AND s.END_POINT = ?;
	`

	rows, err := db.Query(query, token, token, feedID, subscription)
	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	}
	defer rows.Close()

	subscriptions := make([]models.Subscription, 0)

	for rows.Next() {
		var subscription models.Subscription
		err := rows.Scan(&subscription.UserID, &subscription.FeedID, &subscription.EndPoint)
		if err != nil {
			fmt.Println("Failed to scan row: ", err)
			continue
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error occurred while iterating over rows: ", err)
		return nil, err
	}

	return subscriptions, nil

}

func DeleteSubscription(token, subscription string, feedID int) (bool, error) {
	stmt, err := db.Prepare(`
	DELETE s
	FROM subscription AS s
	INNER JOIN user_key AS uk ON uk.USER_ID = s.USER_ID AND uk.USER_KEY = ?
	WHERE s.FEED_ID = ? AND s.END_POINT = ?`)
	if err != nil {
		fmt.Println("Failed to prepare delete subscription statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(token, feedID, subscription)
	if err != nil {
		fmt.Println("Failed to execute delete subscription statement: ", err)
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

	fmt.Println("No rows were affected: ", rowsAffected)
	return false, ErrNotFound
}

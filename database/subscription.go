package database

import (
	"barnyard/api/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertSubscription(token, subscription string, feedID int) (bool, error) {
	stmt, err := db.Prepare(`INSERT INTO subscription (USER_ID, FEED_ID, END_POINT) 
	SELECT fmp.USER_ID, fmp.FEED_ID, ?	
	from feed_member_permission AS fmp
	INNER JOIN user_key as uk ON uk.USER_ID = fmp.USER_ID AND uk.USER_KEY = ?
	WHERE fmp.FEED_ID = ?   ;
	
	`)
	if err != nil {
		fmt.Println("Failed to prepare insert subscription statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(subscription, token, feedID)
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
	SELECT fmp.USER_ID, fmp.FEED_ID, s.END_POINT
	from feed_member_permission AS fmp
	INNER JOIN user_key as uk ON uk.USER_ID = fmp.USER_ID AND uk.USER_KEY = ?
	INNER JOIN subscription as s ON s.FEED_ID = fmp.FEED_ID AND s.USER_ID = fmp.USER_ID
		WHERE fmp.FEED_ID = ? AND s.END_POINT = ?;
	`

	rows, err := db.Query(query, token, feedID, subscription)
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

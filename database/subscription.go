package database

import (
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

/*
func SelectMilestone(token string, feedID int) ([]models.Milestone, error) {

	query := `
	SELECT m.MILESTONE_ID, p.FEED_ID, m.MILESTONE_NAME, m.MILESTONE_DAY, m.MILESTONE_COLOR
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
		INNER JOIN milestone as m ON m.FEED_ID = p.FEED_ID
		WHERE p.FEED_ID = ?;
	`

	rows, err := db.Query(query, token, token, feedID)
	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	}
	defer rows.Close()

	milestones := make([]models.Milestone, 0)

	for rows.Next() {
		var milestone models.Milestone
		err := rows.Scan(&milestone.ID, &milestone.FeedID, &milestone.Name, &milestone.Day, &milestone.Color)
		if err != nil {
			fmt.Println("Failed to scan row: ", err)
			continue
		}
		milestones = append(milestones, milestone)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error occurred while iterating over rows: ", err)
		return nil, err
	}

	return milestones, nil

}

func DeleteMilestone(token string, feedID, milestoneID int) (bool, error) {
	stmt, err := db.Prepare(`
	DELETE m
	FROM milestone AS m
	INNER JOIN feed AS f ON f.FEED_ID = m.FEED_ID
	INNER JOIN user_key AS uk ON uk.USER_ID = f.OWNER_ID AND uk.USER_KEY = ?
	WHERE f.FEED_ID = ? AND m.MILESTONE_ID = ?`)
	if err != nil {
		fmt.Println("Failed to prepare delete milestone statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(token, feedID, milestoneID)
	if err != nil {
		fmt.Println("Failed to execute delete milestone statement: ", err)
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
	return false, nil
}
*/

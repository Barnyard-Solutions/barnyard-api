package database

import (
	"barnyard/api/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertMilestone(token, name, color string, feedID, day int) (bool, error) {
	stmt, err := db.Prepare(`INSERT INTO milestone (FEED_ID, MILESTONE_NAME, MILESTONE_DAY, MILESTONE_COLOR) 
	SELECT fmpd.FEED_ID, ?, ?, ? from feed_member_permission_detail AS fmpd
	INNER JOIN user_key as uk ON uk.USER_ID = fmpd.USER_ID AND uk.USER_KEY = ?
	WHERE fmpd.FEED_ID = ?  AND fmpd.manage_milestone = 1;	
	`)
	if err != nil {
		fmt.Println("Failed to prepare insert milestone statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, day, color, token, feedID)
	if err != nil {
		fmt.Println("Failed to execute insert milestone statement: ", err)
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

func SelectMilestone(token string, feedID int) ([]models.Milestone, error) {

	query := `
	SELECT m.MILESTONE_ID, fmpd.FEED_ID, m.MILESTONE_NAME, m.MILESTONE_DAY, m.MILESTONE_COLOR
	from feed_member_permission_detail AS fmpd
	INNER JOIN user_key as uk ON uk.USER_ID = fmpd.USER_ID AND uk.USER_KEY = ?
	INNER JOIN milestone as m ON m.FEED_ID = fmpd.FEED_ID
	WHERE fmpd.FEED_ID = ?  AND fmpd.manage_milestone = 1;
	`

	rows, err := db.Query(query, token, feedID)
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
	return false, ErrNotFound
}

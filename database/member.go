package database

import (
	"barnyard/api/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertMemeber(token, mail string, feedID, permission int) (bool, error) {
	stmt, err := db.Prepare(`INSERT INTO feed_member (FEED_ID, USER_ID, USER_PERMISSION_LEVEL) 
	SELECT f.FEED_ID, u.USER_ID, ? FROM feed as f
	INNER JOIN user_key as uk ON f.OWNER_ID = uk.USER_ID AND uk.USER_KEY = ?
	INNER JOIN user as u ON UPPER(u.USER_MAIL) = UPPER(?) 
	where f.FEED_ID = ?
	`)
	if err != nil {
		fmt.Println("Failed to prepare insert member statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(permission, token, mail, feedID)
	if err != nil {
		fmt.Println("Failed to execute insert member statement: ", err)
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

func SelectMembers(token string, feedID int) ([]models.Member, error) {
	query := `SELECT fmp.USER_ID, u.USER_MAIL, fmp.permission AS permission
	FROM barnyard.feed_member_permission as fmp
	INNER JOIN barnyard.user as u ON fmp.USER_ID = u.USER_ID  
	INNER JOIN barnyard.user_key as uk ON uk.USER_KEY = ? 
	INNER JOIN barnyard.feed_member_permission as fmp2 ON fmp2.USER_ID = uk.USER_ID AND fmp.FEED_ID = fmp2.FEED_ID
	WHERE fmp.FEED_ID = ?	
	`

	rows, err := db.Query(query, token, feedID)
	if err != nil {
		fmt.Println("Failed to execute query: ", err)
		return nil, err
	}
	defer rows.Close()

	members := make([]models.Member, 0)

	for rows.Next() {
		var member models.Member
		err := rows.Scan(&member.ID, &member.Mail, &member.Permission)
		if err != nil {
			fmt.Println("Failed to scan row: ", err)
			continue
		}
		members = append(members, member)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error occurred while iterating over rows: ", err)
		return nil, err
	}

	return members, nil
}

/*

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
*/

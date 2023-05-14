package database

import (
	"barnyard/api/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertMemeber(token, mail string, feedID, permission int) (bool, error) {
	stmt, err := db.Prepare(`INSERT INTO feed_member (FEED_ID, USER_ID, USER_PERMISSION_LEVEL) 
	SELECT fmpd.FEED_ID, u.USER_ID, ? 
	from feed_member_permission_detail AS fmpd
	INNER JOIN user_key as uk ON uk.USER_ID = fmpd.USER_ID AND uk.USER_KEY = ?
	INNER JOIN user as u ON UPPER(u.USER_MAIL) = UPPER(?) 
	where fmpd.FEED_ID = ? AND fmpd.manage_user = 1;
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

func UpdateMemeber(token, mail string, feedID, permission int) (bool, error) {
	stmt, err := db.Prepare(`UPDATE feed_member as fm
	
	INNER JOIN barnyard.feed_member_permission_detail AS fmpd ON fm.FEED_ID = fmpd.FEED_ID
	INNER JOIN barnyard.user_key as uk ON uk.USER_ID = fmpd.USER_ID AND uk.USER_KEY = ?
	INNER JOIN barnyard.user as u ON UPPER(u.USER_MAIL) = UPPER(?) AND fm.USER_ID = u.USER_ID 
	
	SET USER_PERMISSION_LEVEL = ? 
	WHERE fm.FEED_ID = ? AND fmpd.manage_user = 1;
	`)
	if err != nil {
		fmt.Println("Failed to prepare update member statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(token, mail, permission, feedID)
	if err != nil {
		fmt.Println("Failed to execute update member statement: ", err)
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

func DeleteMember(token, mail string, feedID int) (bool, error) {
	stmt, err := db.Prepare(`
	DELETE fm
	FROM feed_member AS fm
	INNER JOIN user_key as uk ON  uk.USER_KEY = ?
	INNER JOIN feed_member_permission_detail AS fmpd ON uk.USER_ID = fmpd.USER_ID AND fmpd.FEED_ID = fm.FEED_ID
	INNER JOIN user as u ON u.USER_ID = fm.USER_ID AND UPPER(u.USER_MAIL) = UPPER(?) 
	where fm.FEED_ID = ? AND fmpd.manage_user = 1;`)
	if err != nil {
		fmt.Println("Failed to prepare delete feed statement: ", err)
		return false, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(token, mail, feedID)
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

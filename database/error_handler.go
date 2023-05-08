package database

import (
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func IsDuplicateEntryError(err error) bool {
	if sqlErr, ok := err.(*mysql.MySQLError); ok {
		if sqlErr.Number == 1062 {
			return true
		}
	}
	return false
}

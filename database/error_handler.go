package database

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrNotFound     = errors.New("not found")
)

func IsDuplicateEntryError(err error) bool {
	if sqlErr, ok := err.(*mysql.MySQLError); ok {
		if sqlErr.Number == 1062 {
			return true
		}
	}
	return false
}

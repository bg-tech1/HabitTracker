package repository

import "database/sql"

type Db interface {
	SELECT(query string, args ...interface{}) (*sql.Rows, error)
	INSERT(query string, args ...interface{}) error
	UPDATE(query string, args ...interface{}) error
	DELETE(query string, args ...interface{}) error
}

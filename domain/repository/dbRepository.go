package repository

import "database/sql"

type Db interface {
	SELECT(query string) (*sql.Rows, error)
	INSERT(query string) error
	UPDATE(query string) error
	DELETE(query string) error
}

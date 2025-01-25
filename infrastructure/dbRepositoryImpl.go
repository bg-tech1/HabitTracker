package infrastructure

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DbRepositoryImpl struct {
	con *sql.DB
}

func NewDbRepositoryImpl() (*DbRepositoryImpl, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	c, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	err = c.Ping()
	if err != nil {
		return nil, err
	}
	c.Exec("CREATE TABLE IF NOT EXISTS sessions (session_id VARCHAR(255) PRIMARY KEY, user_id VARCHAR(255))")
	return &DbRepositoryImpl{con: c}, nil
}

func (d *DbRepositoryImpl) SELECT(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := d.con.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *DbRepositoryImpl) INSERT(query string, args ...interface{}) error {
	return d.ExecQuery(query, args...)
}

func (d *DbRepositoryImpl) UPDATE(query string, args ...interface{}) error {
	return d.ExecQuery(query, args...)
}

func (d *DbRepositoryImpl) DELETE(query string, args ...interface{}) error {
	return d.ExecQuery(query, args...)
}

func (d *DbRepositoryImpl) ExecQuery(query string, args ...interface{}) error {
	_, err := d.con.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

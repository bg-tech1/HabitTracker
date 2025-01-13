package infrastructure

import (
	"database/sql"
	"fmt"
	"os"
)

type DbRepositoryImpl struct {
	con *sql.DB
}

func NewDbRepositoryImpl() (*DbRepositoryImpl, error) {
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open(os.Getenv("DB_DRIVER_NAME"), dbInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &DbRepositoryImpl{con: db}, nil
}

func (d *DbRepositoryImpl) SELECT(query string) (*sql.Rows, error) {
	rows, err := d.con.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows, nil
}

func (d *DbRepositoryImpl) INSERT(query string) error {
	return d.ExecQuery(query)
}

func (d *DbRepositoryImpl) UPDATE(query string) error {
	return d.ExecQuery(query)
}

func (d *DbRepositoryImpl) DELETE(query string) error {
	return d.ExecQuery(query)
}

func (d *DbRepositoryImpl) ExecQuery(query string) error {
	_, err := d.con.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

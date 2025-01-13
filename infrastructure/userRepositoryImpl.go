package infrastructure

import (
	"fmt"
	"habittracker/domain/repository"
)

type UserRepositoryImpl struct {
	dbCon repository.Db
}

func NewUserRepositoryImpl() (*UserRepositoryImpl, error) {
	db, err := NewDbRepositoryImpl()
	if err != nil {
		return nil, err
	}
	return &UserRepositoryImpl{dbCon: db}, nil
}

func (u *UserRepositoryImpl) LoginUser(id string, password string) (bool, error) {
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM users WHERE id = %s AND password = %s)", id, password)
	rows, err := u.dbCon.SELECT(query)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	var exists bool
	for rows.Next() {
		err := rows.Scan(&exists)
		if err != nil {
			return false, err
		}
	}
	return exists, nil
}

func (u *UserRepositoryImpl) RegisterUser(id string, password string) error {
	query := fmt.Sprintf("INSERT INTO users (id, password) VALUES (%s, %s)", id, password)
	return u.dbCon.INSERT(query)
}

func (u *UserRepositoryImpl) GetUser(id string) (*repository.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = %s", id)
	rows, err := u.dbCon.SELECT(query)
	if err != nil {
		return nil, err
	}
	user := &repository.User{}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(user.Id, user.Password, user.Name)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

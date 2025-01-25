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
	db.con.Exec("CREATE TABLE IF NOT EXISTS users (id VARCHAR(255) PRIMARY KEY, password VARCHAR(255))")
	if err != nil {
		return nil, err
	}
	return &UserRepositoryImpl{dbCon: db}, nil
}

func (u *UserRepositoryImpl) LoginUser(id string) (*repository.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = ($1)")
	rows, err := u.dbCon.SELECT(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user := &repository.User{}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Password)
		if err != nil {
			fmt.Println("パースエラー", err)
			return nil, err
		}
	}
	return user, nil
}

func (u *UserRepositoryImpl) RegisterUser(id string, password string) error {
	query := fmt.Sprintf("INSERT INTO users (id, password) VALUES ($1, $2)")
	return u.dbCon.INSERT(query, id, password)
}

func (u *UserRepositoryImpl) GetUser(id string) (*repository.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE id = ($1)")
	rows, err := u.dbCon.SELECT(query, id)
	if err != nil {
		return nil, err
	}
	user := &repository.User{}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Password)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (u *UserRepositoryImpl) CreateSession(userID string, sessionID string) error {
	query := fmt.Sprintf("INSERT INTO sessions (id, user_id) VALUES ($1, $2)")
	return u.dbCon.INSERT(query, sessionID, userID)
}

func (u *UserRepositoryImpl) GetUserId(sessionID string) (string, error) {
	query := fmt.Sprintf("SELECT user_id FROM sessions WHERE id = ($1)")
	rows, err := u.dbCon.SELECT(query, sessionID)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	var user_id string
	for rows.Next() {
		err := rows.Scan(&user_id)
		if err != nil {
			return "", err
		}
	}
	return user_id, nil
}

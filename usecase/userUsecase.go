package usecase

import (
	"fmt"
	"habittracker/domain/repository"
	"habittracker/infrastructure"
	"habittracker/pkg/util"
)

type UserBlg interface {
	LoginUser(id string, password string, sessionID string) (bool, error)
	RegisterUser(id string, password string) error
}

type UserBlgImpl struct {
	ur repository.UserRepository
}

func NewUserBlgImpl() (*UserBlgImpl, error) {
	ur, err := infrastructure.NewUserRepositoryImpl()
	if err != nil {
		return nil, err
	}
	return &UserBlgImpl{ur: ur}, nil
}

func (ub *UserBlgImpl) LoginUser(id string, password string, sessionID string) (bool, error) {
	user, err := ub.ur.LoginUser(id)
	if err != nil {
		fmt.Println("ログインエラー", err)
		return false, err
	}
	exists := util.ComparePassword(user.Password, password)
	if !exists {
		fmt.Println("パスワードが一致しません")
		return false, nil
	}
	ub.ur.CreateSession(user.Id, sessionID)
	return exists, nil
}

func (ub *UserBlgImpl) RegisterUser(id string, password string) error {
	err := ub.ur.RegisterUser(id, password)
	if err != nil {
		return err
	}
	return nil
}

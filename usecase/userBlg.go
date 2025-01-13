package usecase

import (
	"habittracker/domain/repository"
	"habittracker/infrastructure"
)

type UserBlg interface {
	LoginUser(id string, password string) (bool, error)
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

func (ub *UserBlgImpl) LoginUser(id string, password string) (bool, error) {
	exists, err := ub.ur.LoginUser(id, password)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (ub *UserBlgImpl) RegisterUser(id string, password string) error {
	err := ub.ur.RegisterUser(id, password)
	if err != nil {
		return err
	}
	return nil
}

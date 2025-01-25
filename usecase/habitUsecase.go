package usecase

import (
	"habittracker/domain/repository"
	"habittracker/infrastructure"

	"github.com/google/uuid"
)

type HabitBlg interface {
	ConfirmHabit(id string) (*repository.Habit, error)
	ConfirmAllHabits(sessionId string) ([]*repository.Habit, error)
	CreateHabit(sessionId string, habitName string) error
	DeleteHabit(sessionId string) error
}

type HabitBlgImpl struct {
	hr repository.HabitRepository
	ur repository.UserRepository
}

func NewHabitBlgImpl() (*HabitBlgImpl, error) {
	hr, err := infrastructure.NewHabitRepositoryImpl()
	ur, err := infrastructure.NewUserRepositoryImpl()
	if err != nil {
		return nil, err
	}
	return &HabitBlgImpl{hr: hr, ur: ur}, nil
}

func (blg *HabitBlgImpl) ConfirmHabit(id string) (*repository.Habit, error) {
	h, err := blg.hr.GetHabit(id)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (blg *HabitBlgImpl) ConfirmAllHabits(sessionId string) ([]*repository.Habit, error) {
	// ユーザーidを取得
	userId, err := blg.ur.GetUserId(sessionId)
	if err != nil {
		return nil, err
	}
	h, err := blg.hr.GetAllHabits(userId)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (blg *HabitBlgImpl) CreateHabit(sessionId string, habitName string) error {
	// habitのidを生成
	uuid := uuid.New().String()
	// ユーザーidを取得
	userId, err := blg.ur.GetUserId(sessionId)
	if err != nil {
		return err
	}
	err = blg.hr.CreateHabit(uuid, userId, habitName)
	if err != nil {
		return err
	}
	return nil
}

func (blg *HabitBlgImpl) DeleteHabit(sessionId string) error {
	// ユーザーidを取得
	userId, err := blg.ur.GetUserId(sessionId)
	if err != nil {
		return err
	}
	err = blg.hr.DeleteHabit(userId)
	if err != nil {
		return err
	}
	return nil
}

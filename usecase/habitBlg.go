package usecase

import (
	"habittracker/domain/repository"
	"habittracker/infrastructure"

	"github.com/google/uuid"
)

type HabitBlg interface {
	ConfirmHabit(id string) (*repository.Habit, error)
	ConfirmAllHabits(userID string) ([]*repository.Habit, error)
	CreateHabit(userID string, habitName string) error
	DeleteHabit(id string) error
}

type HabitBlgImpl struct {
	hr repository.HabitRepository
}

func NewHabitBlgImpl() (*HabitBlgImpl, error) {
	hr, err := infrastructure.NewHabitRepositoryImpl()
	if err != nil {
		return nil, err
	}
	return &HabitBlgImpl{hr: hr}, nil
}

func (blg *HabitBlgImpl) ConfirmHabit(id string) (*repository.Habit, error) {
	h, err := blg.hr.GetHabit(id)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (blg *HabitBlgImpl) ConfirmAllHabits(userID string) ([]*repository.Habit, error) {
	h, err := blg.hr.GetAllHabits(userID)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (blg *HabitBlgImpl) CreateHabit(userID string, habitName string) error {
	uuid := uuid.New().String()
	err := blg.hr.CreateHabit(uuid, userID, habitName)
	if err != nil {
		return err
	}
	return nil
}

func (blg *HabitBlgImpl) DeleteHabit(id string) error {
	err := blg.hr.DeleteHabit(id)
	if err != nil {
		return err
	}
	return nil
}

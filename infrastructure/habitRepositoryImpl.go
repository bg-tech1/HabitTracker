package infrastructure

import (
	"fmt"
	"habittracker/domain/repository"
)

type HabitRepositoryImpl struct {
	dbCon repository.Db
}

func NewHabitRepositoryImpl() (*HabitRepositoryImpl, error) {
	db, err := NewDbRepositoryImpl()
	if err != nil {
		return nil, err
	}
	return &HabitRepositoryImpl{dbCon: db}, nil
}

func (h *HabitRepositoryImpl) CreateHabit(habitID string, userID string, habitName string) error {
	query := fmt.Sprintf("INSERT INTO habits (id, user_id, habit_name) VALUES (%s, %s, %s)", habitID, userID, habitName)
	return h.dbCon.INSERT(query)
}

func (h *HabitRepositoryImpl) DeleteHabit(habitID string) error {
	query := fmt.Sprintf("DELETE FROM habits WHERE id = %s", habitID)
	return h.dbCon.DELETE(query)
}

func (h *HabitRepositoryImpl) GetHabit(id string) (*repository.Habit, error) {
	query := fmt.Sprintf("SELECT * FROM habits WHERE id = %s", id)
	rows, err := h.dbCon.SELECT(query)
	if err != nil {
		return nil, err
	}
	habit := &repository.Habit{}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(habit.Id, habit.UserId, habit.HabitName, habit.RecordId)
		if err != nil {
			return nil, err
		}
	}
	return habit, nil
}

func (h *HabitRepositoryImpl) GetAllHabits(userID string) ([]*repository.Habit, error) {
	query := fmt.Sprintf("SELECT * FROM habits WHERE user_id = %s", userID)
	rows, err := h.dbCon.SELECT(query)
	if err != nil {
		return nil, err
	}
	habits := []*repository.Habit{}
	defer rows.Close()
	for rows.Next() {
		habit := &repository.Habit{}
		err := rows.Scan(habit.Id, habit.UserId, habit.HabitName, habit.RecordId)
		if err != nil {
			return nil, err
		}
		habits = append(habits, habit)
	}
	return habits, nil
}

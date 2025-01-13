package repository

import "time"

type Record struct {
	Id     string
	Date   time.Time
	IsDone bool
}

type Habit struct {
	Id        string
	UserId    string
	HabitName string
	RecordId  string
}

type HabitRepository interface {
	CreateHabit(habitID string, userID string, habitName string) error
	DeleteHabit(habitID string) error
	GetHabit(habitID string) (*Habit, error)
	GetAllHabits(userID string) ([]*Habit, error)
	//TODO: Add UpdateHabit
}

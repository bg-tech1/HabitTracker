package repository

import "time"

type Record struct {
	Id     string    `json:"id"`
	Date   time.Time `json:"date"`
	IsDone bool      `json:"is_done"`
}

type Habit struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	HabitName string `json:"habit_name"`
	RecordId  string `json:"record_id"`
}

type HabitRepository interface {
	CreateHabit(habitID string, userID string, habitName string) error
	DeleteHabit(habitID string) error
	GetHabit(habitID string) (*Habit, error)
	GetAllHabits(sessionID string) ([]*Habit, error)
	//TODO: Add UpdateHabit
}

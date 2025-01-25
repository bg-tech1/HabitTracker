package infrastructure

import (
	"database/sql"
	"fmt"
	"habittracker/domain/repository"
)

type HabitRepositoryImpl struct {
	dbCon repository.Db
}

func NewHabitRepositoryImpl() (*HabitRepositoryImpl, error) {
	db, err := NewDbRepositoryImpl()
	db.con.Exec("CREATE TABLE IF NOT EXISTS habits (id VARCHAR(255) PRIMARY KEY, user_id VARCHAR(255), habit_name VARCHAR(255), record_id VARCHAR(255))")
	if err != nil {
		return nil, err
	}
	return &HabitRepositoryImpl{dbCon: db}, nil
}

func (h *HabitRepositoryImpl) CreateHabit(habitID string, userID string, habitName string) error {
	query := fmt.Sprintf("INSERT INTO habits (id, user_id, habit_name) VALUES ($1, $2, $3)")
	return h.dbCon.INSERT(query, habitID, userID, habitName)
}

func (h *HabitRepositoryImpl) DeleteHabit(habitID string) error {
	query := fmt.Sprintf("DELETE FROM habits WHERE id = ($1)")
	return h.dbCon.DELETE(query, habitID)
}

func (h *HabitRepositoryImpl) GetHabit(id string) (*repository.Habit, error) {
	query := fmt.Sprintf("SELECT * FROM habits WHERE id = ($1)")
	rows, err := h.dbCon.SELECT(query, id)
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

func (h *HabitRepositoryImpl) GetAllHabits(userId string) ([]*repository.Habit, error) {
	query := "SELECT * FROM habits WHERE user_id = $1"
	rows, err := h.dbCon.SELECT(query, userId)
	if err != nil {
		return nil, err
	}
	habits := []*repository.Habit{}
	defer rows.Close()
	for rows.Next() {
		habit := &repository.Habit{}
		var recordId sql.NullString
		err := rows.Scan(&habit.Id, &habit.UserId, &habit.HabitName, &recordId)
		if err != nil {
			return nil, err
		}
		if recordId.Valid {
			habit.RecordId = recordId.String
		} else {
			habit.RecordId = ""
		}
		habits = append(habits, habit)
	}
	return habits, nil
}

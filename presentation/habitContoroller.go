package presentation

import (
	"habittracker/usecase"

	"github.com/gin-gonic/gin"
)

type HabitController interface {
	ConfirmHabit(c *gin.Context)
	CreateHabit(c *gin.Context)
	DeleteHabit(c *gin.Context)
}

type HabitControllerImpl struct {
	hr usecase.HabitBlg
	ur usecase.UserBlg
}

func NewHabitControllerImpl() *HabitControllerImpl {
	hr, err := usecase.NewHabitBlgImpl()
	if err != nil {
		panic(err)
	}
	ur, err := usecase.NewUserBlgImpl()
	if err != nil {
		panic(err)
	}
	return &HabitControllerImpl{hr: hr, ur: ur}
}

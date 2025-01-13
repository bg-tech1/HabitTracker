package presentation

import (
	"habittracker/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hc *HabitControllerImpl) CreateHabit(c *gin.Context) {
	habit := repository.Habit{}
	if err := c.ShouldBindJSON(&habit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		return
	}
	err := hc.hr.CreateHabit(habit.UserId, habit.HabitName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Create Success": "true"})
}

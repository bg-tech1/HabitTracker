package presentation

import (
	"habittracker/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hc *HabitControllerImpl) ConfirmHabit(c *gin.Context) {
	h := repository.Habit{}
	if err := c.ShouldBindJSON(&h); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		return
	}
	habit, err := hc.hr.ConfirmHabit(h.Id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Habit": habit})
}

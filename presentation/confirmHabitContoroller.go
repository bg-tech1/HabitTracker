package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hc *HabitControllerImpl) ConfirmHabit(c *gin.Context) {
	habit, err := hc.hr.ConfirmHabit(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Habit": habit})
}

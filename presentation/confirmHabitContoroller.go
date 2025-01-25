package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (hc *HabitControllerImpl) ConfirmHabit(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Can't get session_id from cookie": err.Error()})
		return
	}
	habit, err := hc.hr.ConfirmAllHabits(sessionID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Can't get habit": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"habit": habit})
}

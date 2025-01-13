package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *HabitControllerImpl) DeleteHabit(c *gin.Context) {
	err := h.hr.DeleteHabit(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Delete Success": "true"})
}

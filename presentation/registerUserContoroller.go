package presentation

import (
	"habittracker/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserControllerImpl) RegisterUser(c *gin.Context) {
	user := repository.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
	}
	err := u.blg.RegisterUser(user.Id, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Register Success": "true"})
}

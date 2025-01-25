package presentation

import (
	"habittracker/domain/repository"
	"habittracker/pkg/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserControllerImpl) RegisterUser(c *gin.Context) {
	user := repository.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		return
	}
	hp, err := util.HashPassword(user.Password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process password"})
		return
	}
	err = u.blg.RegisterUser(user.Id, string(hp))
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

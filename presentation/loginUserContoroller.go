package presentation

import (
	"habittracker/domain/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserControllerImpl) LoginUser(c *gin.Context) {
	user := repository.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		return
	}
	exists, err := uc.blg.LoginUser(user.Id, user.Password)
	//Blgの結果に応じてレスポンスを返す
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		return
	}
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Login Success": "true"})
}

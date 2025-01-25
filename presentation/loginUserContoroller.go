package presentation

import (
	"habittracker/domain/repository"
	"habittracker/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserControllerImpl) LoginUser(c *gin.Context) {
	user := repository.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Bad Request!!": err})
		return
	}

	// TODO:sesisonの管理については後で考える
	// cookieの保存
	sessionID := util.GenerateSessionID()
	c.SetCookie("session_id", sessionID, 3600, "/", "", false, false)
	exists, err := uc.blg.LoginUser(user.Id, user.Password, sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err})
		return
	}
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"Unauthorized": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Login Success": "true", "redirect": "/view/dashboard.html"})
}

package presentation

import (
	"habittracker/usecase"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	LoginUser(c *gin.Context)
	RegisterUser(c *gin.Context)
}

type UserControllerImpl struct {
	blg usecase.UserBlg
}

// TODO:usecase側ではinterfaceで受け取り、メソッドを呼び出す
func NewUserControllerImpl() *UserControllerImpl {
	blg, err := usecase.NewUserBlgImpl()
	if err != nil {
		panic(err)
	}
	return &UserControllerImpl{blg: blg}
}

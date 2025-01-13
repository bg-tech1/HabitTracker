package main

import (
	"habittracker/pkg"
	"habittracker/presentation"

	"github.com/gin-gonic/gin"
)

var uc = presentation.NewUserControllerImpl()

var hc = presentation.NewHabitControllerImpl()

func main() {
	engine := gin.Default()
	// ロギング
	engine.Use(pkg.Logger)
	appEngine := engine.Group("/app")
	{
		h := appEngine.Group("/habit")
		{
			h.POST("/confirm", hc.ConfirmHabit)
			h.POST("/delete", hc.DeleteHabit)
			h.POST("/create", hc.CreateHabit)
		}
		u := appEngine.Group("/user")
		{
			u.POST("/login", uc.LoginUser)
			u.POST("/register", uc.RegisterUser)
		}
	}
	engine.Run(":8080")
}

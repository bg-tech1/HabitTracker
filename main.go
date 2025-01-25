package main

import (
	"fmt"
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
	// 静的ファイルの配信
	engine.Static("/view", "./view")
	appEngine := engine.Group("/app")
	{
		h := appEngine.Group("/habit")
		{
			h.GET("/dashboard", hc.ConfirmHabit)
			h.POST("/delete", hc.DeleteHabit)
			h.POST("/create", hc.CreateHabit)
		}
		u := appEngine.Group("/user")
		{
			u.POST("/login", uc.LoginUser)
			u.POST("/register", uc.RegisterUser)
			// u.GET("/info", uc.GetUserInfo)
		}
	}
	engine.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

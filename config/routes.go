package config

import (
	"app/modules/auth"
	"app/modules/home"
	"app/modules/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(e *gin.Engine) {
	e.GET("/login", auth.HandleLogin)
	e.POST("/login", auth.HandleLoginAction)
	e.GET("/logout", auth.HandleLogout)
	e.GET("/register", user.HandleRegister)
	e.GET("/home", home.HandleHome)
}

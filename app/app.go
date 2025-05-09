package app

import (
	"project/app/auth"
	"project/app/home"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	SessionSecret = "617b31b60b908a16fce4ef735c93e3586da7f0d1ce853875df85e1ceb3b4c4b8"
	SessionName   = "session"
	TemplatesPath = "templates/**/*"
	StaticPath    = "static"
)

func SetupApp() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob(TemplatesPath)
	r.Static(StaticPath, StaticPath)

	// Middlewares
	r.Use(sessions.Sessions(SessionName, cookie.NewStore([]byte(SessionSecret))))
	r.Use(auth.AuthMiddleware)

	// Routes
	r.GET("/login", auth.LoginRoute)
	r.POST("/login", auth.LoginActionRoute)
	r.GET("/logout", auth.LogoutRoute)
	r.GET("/home", home.HomeRoute)

	return r
}

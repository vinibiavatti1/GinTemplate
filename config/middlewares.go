package config

import (
	"app/modules/auth"
	"app/session"

	"github.com/gin-gonic/gin"
)

func SetupMiddlewares(e *gin.Engine) {
	e.Use(session.SessionMiddleware())
	e.Use(auth.AuthMiddleware)
}

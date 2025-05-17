package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	SessionSecret = "617b31b60b908a16fce4ef735c93e3586da7f0d1ce853875df85e1ceb3b4c4b8"
	SessionName   = "session"
)

func SessionMiddleware() gin.HandlerFunc {
	return sessions.Sessions(SessionName, cookie.NewStore([]byte(SessionSecret)))
}

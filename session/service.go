package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const SessionKey = "Session"

func StartSession(c *gin.Context, data *Session) {
	s := sessions.Default(c)
	s.Set(SessionKey, data)
	s.Save()
}

func TerminateSession(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
}

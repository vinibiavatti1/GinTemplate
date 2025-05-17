package template

import (
	"app/modules/message"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	session := sessions.Default(c)
	data["Message"] = message.GetFromContext(c)
	data["Session"] = session.Get("Session")
	c.HTML(code, name, data)
}

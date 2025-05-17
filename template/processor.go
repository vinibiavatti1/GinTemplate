package template

import (
	"app/modules/message"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Render(c *gin.Context, code int, name string, data gin.H) {
	data = injectCommonData(c, data)
	c.HTML(code, name, data)
}

func injectCommonData(c *gin.Context, data gin.H) gin.H {
	session := sessions.Default(c)
	data["Message"] = message.GetFromContext(c)
	data["Session"] = session.Get("Session")
	return data
}

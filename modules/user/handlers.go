package user

import (
	"app/modules/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRegister(c *gin.Context) {
	template.Render(c, http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

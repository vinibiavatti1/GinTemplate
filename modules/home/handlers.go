package home

import (
	"app/modules/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHome(c *gin.Context) {
	template.Render(c, http.StatusOK, "home.html", gin.H{
		"title": "Hello World!",
	})
}

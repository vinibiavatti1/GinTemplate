package auth

import (
	"app/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleLogin(c *gin.Context) {
	template.Render(c, http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func HandleLogout(c *gin.Context) {
	Logout(c)
	c.Redirect(http.StatusOK, "/login?status=logoutSuccess")
}

func HandleLoginAction(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	ok, err := Login(c, email, password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.Redirect(http.StatusOK, "/login?status=loginFailed")
		return
	}
	c.Redirect(http.StatusOK, "/login?status=loginSuccess")
}

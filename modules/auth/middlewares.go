package auth

import "github.com/gin-gonic/gin"

func AuthMiddleware(c *gin.Context) {
	c.Next()
}

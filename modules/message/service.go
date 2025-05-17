package message

import "github.com/gin-gonic/gin"

const MessageParam = "message"

func GetFromContext(c *gin.Context) *Message {
	return Messages[c.Query(MessageParam)]
}

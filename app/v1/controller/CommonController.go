package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetUserIdFromSession(c *gin.Context) int {
	session := sessions.Default(c)
	if id := session.Get("user_id"); id != nil {
		return id.(int)
	}
	return -1
}

package middleware

import (
	"Todo/app/helper"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authority() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("user_id")
		if id == nil {
			c.JSON(http.StatusOK, helper.ApiReturn(http.StatusUnauthorized, "权限不足,请先登录"))
			c.Abort()
		}
	}
}

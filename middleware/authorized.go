package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"network/global/session"
)

func Authorized(c *gin.Context)  {
	if !session.IsLogin(c) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}
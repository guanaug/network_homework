package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"network/global/logger"
	"network/global/session"
	"network/model/user"
)

func Authorized(c *gin.Context) {
	if !session.IsLogin(c) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

var admin string
func AdminAuthorized(c *gin.Context) {
	username := session.User(c)
	if len(username) == 0 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// 由于admin账号不可变，故如果已经初始化过admin，则不再查询数据库
	if len(admin) == 0 {
		u := &user.User{Account:username}
		if ok, err := u.IsAdmin(); err != nil || ok {
			if err != nil {
				logger.Logger().Warn("query administrator err:", err)
			}
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		admin = u.Account
	}

	if admin != username {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

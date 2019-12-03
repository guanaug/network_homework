package session

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"network/global/constant"
	"network/global/logger"
	"network/global/session"
	"network/model/user"
	"network/util/password"
	"strings"
)

// TODO 可以考虑添加验证码
// Account校验
// 密码校验
func Login(c *gin.Context) {
	loginInfo := struct {
		Account  string `binding:"required"`
		Password string `binding:"required"`
	}{}

	if err := c.BindJSON(&loginInfo); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "账号和密码不能为空！"})
		return
	}

	if session.IsLogin(c) {
		c.Status(http.StatusOK)
		return
	}

	u := &user.User{
		Account:  strings.TrimSpace(loginInfo.Account),
		Password: password.New(strings.TrimSpace(loginInfo.Password)),
	}

	ok, err := u.Login()
	if err != nil {
		logger.Logger().Warn("query login info error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "账号或密码错误！"})
		return
	}

	if err := session.Login(c, loginInfo.Account); err != nil {
		logger.Logger().Warn("add session error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	userBriefInfo := struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}{ID: u.ID, Name: u.Name}

	// 记录用户ID，方便日志拦截器使用
	c.Set(constant.KeyUserID, u.ID)

	c.JSON(http.StatusOK, &userBriefInfo)
}

func Logout(c *gin.Context) {
	if err := session.Logout(c); err != nil {
		logger.Logger().Warn("delete session error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

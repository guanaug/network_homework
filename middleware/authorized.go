package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"network/global/constant"
	"network/global/logger"
	"network/global/session"
	"network/model/department"
	"network/model/user"
)

func Authorized(c *gin.Context) {
	if !session.IsLogin(c) {
		logger.Logger().Debug("未登录")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

var admin string

func AdminAuthorized(c *gin.Context) {
	// 由于前面已经检查是否登录，故这里username长度必然大于0
	username := session.GetUser(c)

	// 由于admin账号不可变，故如果已经初始化过admin，则不再查询数据库
	if len(admin) == 0 {
		u := &user.User{Account: username}
		if ok, err := u.IsAdmin(); err != nil || !ok {
			if err != nil {
				logger.Logger().Warn("query administrator err:", err)
			}
			logger.Logger().Debug(u.Account, "没有管理员权限")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		admin = u.Account
	}

	if admin != username {
		logger.Logger().Debug(username, "不是管理员")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func CityAndDistrictAuthorized(c *gin.Context) {
	// 由于前面已经检查是否登录，故这里username长度必然大于0
	username := session.GetUser(c)

	u, err := user.OneByAccount(username)
	if err != nil {
		logger.Logger().Warn("get user error:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	depart := department.Department{ID: u.Department}
	if ok, err := depart.IsRoleOr(constant.TypeUserCity, constant.TypeUserDistrict); err != nil || !ok {
		if err != nil {
			logger.Logger().Warn("query administrator err:", err)
		}
		logger.Logger().Debug(u.Account, "没有", constant.GetDepartmentTypeName(constant.TypeUserCity), "和",
			constant.GetDepartmentTypeName(constant.TypeUserDistrict), "权限")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

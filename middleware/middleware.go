package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"network/global/constant"
	"network/global/logger"
	"network/model/loginlog"
)

func UserLoginLog(c *gin.Context)  {
	c.Next()

	ip := c.ClientIP()
	id := c.GetInt64(constant.KeyUserID)
	if 0 == id {
		// id 为 0 情况，说明已经登录过，不写日志，直接返回
		return
	}

	ul := loginlog.UserLog{
		UserID:    id,
		IP:        ip,
	}
	if err := ul.Insert(); err != nil {
		// 如果写入登录日志失败，为了不影响主流程，只记录错误，不中断
		logger.Logger().Warn("写入登录日志失败:", err)
		return
	}
}

// TODO 这只是临时解决办法，务必找到较完美方案
func AllowCORS(c *gin.Context)  {
	c.Next()
	if logger.Logger().Level == logrus.DebugLevel {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	}
}
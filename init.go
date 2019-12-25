package main

import (
	"github.com/gin-gonic/gin"
	"network/controller/department"
	"network/controller/session"
	"network/controller/transaction"
	"network/controller/user"
	"network/middleware"
)

func routerRegister(router *gin.Engine) {
	authorized := router.Group("/")
	{
		authorized.Use(middleware.Authorized)

		groupUser := authorized.Group("/user")
		{
			groupUser.POST("/type", user.ListWithType)
			{
				groupUser.Use(middleware.AdminAuthorized)
				// 添加用户
				groupUser.POST("", user.Add)
				// 删除用户
				groupUser.DELETE("/:id", user.Delete)
				// 修改用户
				groupUser.PUT("", user.Modify)
				// 获取用户详细信息
				groupUser.GET("/:id", user.Info)
				// 获取用户列表
				groupUser.GET("", user.List)
			}
		}

		groupDepartment := authorized.Group("/department")
		{
			// 获取部门列表
			groupDepartment.GET("", department.List)
			groupDepartment.POST("/type", department.ListWithType)
			{
				groupDepartment.Use(middleware.AdminAuthorized)
				// 添加部门
				groupDepartment.POST("", department.Add)
				// 修改部门
				groupDepartment.PUT("", department.Modify)
				// 删除部门
				groupDepartment.DELETE("/:id", department.Delete)
				// 获取部门详细信息
				groupDepartment.GET("/:id", department.Info)
			}
		}

		groupTransaction := authorized.Group("/transaction")
		{
			// 获取事件信息
			groupTransaction.GET("", transaction.List)
			// 修改事件
			groupTransaction.PUT("", transaction.Modified)
			{
				groupTransaction.Use(middleware.CityAndDistrictAuthorized)
				// 添加事件
				groupTransaction.POST("", transaction.Add)
				// 事件统计
				groupTransaction.GET("/statistic", transaction.Statistic)
				groupTransaction.GET("/statistic/array", transaction.StatisticArray)
			}
		}

		// 注销登录
		authorized.DELETE("/session", session.Logout)
		//查看登录日志
		authorized.GET("/session/log", middleware.AdminAuthorized, session.LoginLog)
	}

	unauthorized := router.Group("/")
	{
		// 用户登录
		unauthorized.POST("/session", middleware.UserLoginLog, session.Login)
	}
}

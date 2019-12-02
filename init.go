package main

import (
	"github.com/gin-gonic/gin"
	"network/controller/department"
	"network/controller/user"
	"network/middleware"
)

func routerRegister(router *gin.Engine) {
	authorized := router.Group("/")
	{
		authorized.Use(middleware.Authorized)
		// 注销登录
		authorized.DELETE("/user", user.Logout)

		groupUser := authorized.Group("/user")
		{
			// 添加用户
			groupUser.POST("/", user.Add)
			// 删除用户
			groupUser.DELETE("/:id", user.Delete)
			// 修改用户
			groupUser.PUT("/", user.Modify)
			// 获取用户列表
			groupUser.GET("/", user.List)
			// 获取用户详细信息
			groupUser.GET("/:id", user.Info)
		}

		groupDepartment := authorized.Group("/department")
		{
			{
				groupDepartment.Use(middleware.AdminAuthorized)
				// 添加部门
				groupDepartment.POST("/", department.Add)
				// 修改部门
				groupDepartment.PUT("/", department.Modify)
				// 删除部门
				groupDepartment.DELETE("/:id", department.Delete)
				// 获取部门列表
				groupDepartment.GET("/", department.List)
				// 获取部门详细信息
				groupDepartment.GET("/:id", department.Info)
			}
		}
	}

	unauthorized := router.Group("/")
	{
		// 用户登录
		unauthorized.POST("/login", user.Login)
	}
}

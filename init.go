package main

import (
	"github.com/gin-gonic/gin"
	"network/controller/department"
	"network/controller/user"
	"network/middleware"
)

func routerRegister(router *gin.Engine)  {
	authorized := router.Group("/")
	{
		authorized.Use(middleware.Authorized)

		{
			// 注销登录
			authorized.DELETE("/user", user.Logout)
			// 添加用户
			authorized.POST("/user", user.Add)
		}

		groupDepartment := authorized.Group("/department")
		{
			{
				// TODO 必须管理员才有以下权限
				//groupDepartment.Use(middleware.AdminAuthorized)
				// 添加部门
				groupDepartment.POST("/", department.Add)
				// 修改部门
				groupDepartment.PUT("/", department.Modified)
				// 删除部门
				groupDepartment.DELETE("/:id", department.Delete)
			}
			// 获取部门列表
			groupDepartment.GET("/", department.List)
			// 获取部门详细信息
			groupDepartment.GET("/:id", department.Info)
		}
	}

	unauthorized := router.Group("/")
	{
		// 用户登录
		unauthorized.POST("/login", user.Login)
	}
}
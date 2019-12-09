package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"network/util/verification"
)

func main() {
	router := gin.Default()

	if err := verification.Register(); err != nil {
		log.Fatal("校验器注册失败:", err)
	}

	routerRegister(router)

	// 临时解决跨域问题
	router.Static("/static", `H:\www\network_homework`)

	if err := router.Run(":10086"); err != nil {
		log.Fatal(err)
	}
}

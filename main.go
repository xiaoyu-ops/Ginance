package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 路由引擎
	r := gin.Default()

	// 定义一个get接口
	// 路径是/api/v1/ping
	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping",
			"status":  "success",
		})
	})

	// 启动服务，监听8080端口
	r.Run(":8080")
}

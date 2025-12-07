package main

import (
	"giance/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1.加载.env 配置
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2.初始化数据库连接

	db.InitDB()

	// 3.启动 gin服务

	r := gin.Default()

	// 定义一个get接口
	// 路径是/api/v1/ping
	r.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// 启动服务，监听8080端口
	r.Run(":8080")
}

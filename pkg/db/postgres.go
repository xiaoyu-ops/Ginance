package db

import (
	"fmt"
	"log"
	"os"

	"giance/internal/model"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

// 全局化DB变量，方便其他包使用
var DB *gorm.DB

//InitDB 初始化数据库连接

func InitDB() {
	// 1.组装连接字符串(DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// 2.连接数据库
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败，检查.env配置", err)
	}

	log.Println("数据库连接成功")

	// 3. 自动迁移
	// 在数据库中自动创建users表
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("数据库迁移失败", err)
	}

	log.Println("数据库表结构迁移成功")
}

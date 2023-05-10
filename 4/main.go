package main

import (
	"4/api"
	"4/common"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	// 初始化Redis客户端
	common.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 测试Redis连接
	_, err := common.RedisClient.Ping(common.RedisClient.Context()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}

	// 初始化Gin路由
	router := gin.Default()

	// 注册路由处理函数
	router.POST("/register", api.RegisterHandler)
	router.POST("/login", api.LoginHandler)
	router.POST("/change-password", api.ChangePasswordHandler)
	router.POST("/reset-password", api.ResetPasswordHandler)

	// 启动服务器
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

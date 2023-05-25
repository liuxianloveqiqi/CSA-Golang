package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
)

var rdb *redis.Client

func main() {
	// 初始化Redis连接
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果有密码，填写密码
		DB:       0,  // 选择默认数据库
	})

	// 测试Redis连接
	pong, err := rdb.Ping(rdb.Context()).Result()
	if err != nil {
		log.Fatal("不能连接到redis:", err)
	}
	fmt.Println("连接redis成功", pong)

	// 创建Gin路由
	router := gin.Default()

	// 定义点赞接口
	router.POST("/like/:likerID/:likedID", likeHandler)

	// 启动服务器
	router.Run(":8080")
}

func likeHandler(c *gin.Context) {
	// 使用redis的哈希，点赞者和被点赞者作 key1 key2 ,是否点赞作value2
	likerID := c.Param("likerID")
	likedID := c.Param("likedID")

	// 检查点赞者是否已经给该被点赞者点过赞
	exists, err := rdb.HExists(rdb.Context(), likerID, likedID).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法检查点赞"})
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "已经点赞过了无法继续点赞"})
		return
	}

	// 在Redis中设置点赞记录，使用1表示已经点赞了
	err = rdb.HSet(rdb.Context(), likerID, likedID, "1").Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法设置点赞状态"})
		return
	}

	// 返回成功信息
	c.JSON(http.StatusOK, gin.H{"message": "点赞成功"})
}

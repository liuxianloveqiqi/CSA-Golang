package api

import (
	"4/common"
	"4/model"
	"crypto/rand"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"math/big"
	"net/http"
)

func LoginHandler(c *gin.Context) {
	// 解析请求体中的JSON数据
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求"})
		return
	}

	// 从Redis中获取密码
	password, err := common.RedisClient.Get(common.RedisClient.Context(), user.Username).Result()
	if err != nil {
		if err == redis.Nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名不存在"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		}
		return
	}

	// 验证密码是否匹配
	if password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码不正确"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}

func RegisterHandler(c *gin.Context) {
	// 解析请求体中的JSON数据
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求"})
		return
	}

	// 检查用户名是否已存在
	exists, err := common.RedisClient.Exists(common.RedisClient.Context(), user.Username).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		return
	}

	if exists == 1 {
		c.JSON(http.StatusConflict, gin.H{"error": "用户名已存在"})
		return
	}

	// 将用户名和密码存储到Redis中
	err = common.RedisClient.Set(common.RedisClient.Context(), user.Username, user.Password, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func ChangePasswordHandler(c *gin.Context) {
	// 解析请求体中的JSON数据
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求"})
		return
	}

	// 检查用户名是否存在
	exists, err := common.RedisClient.Exists(common.RedisClient.Context(), user.Username).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		return
	}

	if exists == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户名不存在"})
		return
	}

	// 更新密码
	err = common.RedisClient.Set(common.RedisClient.Context(), user.Username, user.Password, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "密码修改成功"})
}

func ResetPasswordHandler(c *gin.Context) {
	// 解析请求体中的JSON数据
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求"})
		return
	}

	// 检查用户名是否存在
	exists, err := common.RedisClient.Exists(common.RedisClient.Context(), user.Username).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		return
	}

	if exists == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户名不存在"})
		return
	}

	// 生成新的随机密码
	newPassword := GeneratePassword(9)

	// 更新密码
	err = common.RedisClient.Set(common.RedisClient.Context(), user.Username, newPassword, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		return
	}

	// 返回新密码
	c.JSON(http.StatusOK, gin.H{"message": "密码已重置", "new_password": newPassword})
}

func GeneratePassword(length int) string {
	// 定义密码包含的字符集
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~!@#$%^&*()_+`-={}|[]\\:\";'<>,.?/"

	// 定义密码长度

	passwordLength := length

	// 初始化密码切片
	password := make([]byte, passwordLength)

	// 生成随机密码
	for i := 0; i < passwordLength; i++ {
		charIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		password[i] = charset[charIndex.Int64()]
	}

	return string(password)
}

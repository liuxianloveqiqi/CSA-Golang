package main

import (
	models "5/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

var DB *gorm.DB

func (s *Server) Initialize() {
	dsn := "root:123456@tcp(localhost:3306)/qq?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	s.DB = db
	s.Router = gin.Default()
	s.InitializeRoutes()
}

func (s *Server) InitializeRoutes() {
	// 登录
	s.Router.POST("/login", s.LoginHandler)

	// 创建账号
	s.Router.POST("/create-account", s.CreateAccountHandler)

	// 加好友
	s.Router.POST("/add-friend", s.AddFriendHandler)

	// 删好友
	s.Router.POST("/delete-friend", s.DeleteFriendHandler)
}

// 处理登录逻辑
func (s *Server) LoginHandler(c *gin.Context) {
	// 解析请求参数
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shouldbind错误"})
		return
	}

	// 查询数据库验证登录信息
	var user models.User
	result := s.DB.Where("username = ?", loginData.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或者密码错误"})
		return
	}
	if user.Password != loginData.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或者密码错误"})
		return
	}

	// 登录成功，返回用户信息或令牌等
	c.JSON(http.StatusOK, gin.H{"message": "登陆成功", "user": user})
}

// 处理创建账号逻辑
func (s *Server) CreateAccountHandler(c *gin.Context) {
	// 解析请求参数
	var userData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shouldbind解析错误"})
		return
	}

	// 检查用户名是否已存在
	var existingUser models.User
	result := s.DB.Where("username = ?", userData.Username).First(&existingUser)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "用户名已存在"})
		return
	}

	// 创建用户账号
	newUser := models.User{
		Username: userData.Username,
		Password: userData.Password,
	}
	result = s.DB.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建用户失败"})
		return
	}

	// 返回创建成功的用户信息
	c.JSON(http.StatusOK, gin.H{"message": "创建用户成功", "user": newUser})
}

// 处理添加好友逻辑
func (s *Server) AddFriendHandler(c *gin.Context) {

	// 解析请求参数
	var friendData struct {
		UserID   uint `json:"user_id"`
		FriendID uint `json:"friend_id"`
	}
	if err := c.ShouldBindJSON(&friendData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shouldbind解析错误"})
		return
	}

	// 检查用户是否存在
	var user models.User
	result := s.DB.First(&user, friendData.UserID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有这个user_ID"})
		return
	}

	// 检查好友是否存在
	var friend models.User
	result = s.DB.First(&friend, friendData.FriendID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "没有这个friend_ID"})
		return
	}

	// 检查是否已经是好友关系
	var existingFriend models.Friend
	result = s.DB.Where("user_id = ? AND friend_id = ?", friendData.UserID, friendData.FriendID).First(&existingFriend)
	if result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "已经是好友了"})
		return
	}
	newFriend := models.Friend{
		UserID:   friendData.UserID,
		FriendID: friendData.FriendID,
	}
	result = s.DB.Create(&newFriend)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "添加好友失败"})
		return
	}

	// 返回添加好友成功的信息
	c.JSON(http.StatusOK, gin.H{"message": "添加好友成功", "friend": newFriend})
}

func (s *Server) DeleteFriendHandler(c *gin.Context) {
	// 解析请求参数
	var friendData struct {
		UserID   uint `json:"user_id"`
		FriendID uint `json:"friend_id"`
	}
	if err := c.ShouldBindJSON(&friendData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shouldbind解析错误"})
		return
	}

	// 删除好友关系
	result := s.DB.Where("user_id = ? AND friend_id = ?", friendData.UserID, friendData.FriendID).Delete(&models.Friend{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除好友失败"})
		return
	}

	// 返回删除成功的信息
	c.JSON(http.StatusOK, gin.H{"message": "删除好友成功"})
}

func main() {
	server := Server{}
	server.Initialize()
	server.Router.Run(":8080")
}

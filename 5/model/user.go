package models

import (
	"gorm.io/gorm"
)

type User struct {
	Username string
	Password string
}

type Friend struct {
	gorm.Model
	UserID   uint
	FriendID uint
}

package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Blog []Blog `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blog"`
  }

type Blog struct {
	gorm.Model
	Title   string `json:"title" form:"form"`
	Content string `json:"content" form:"content"`
	UserID int `json:"user_id"`
	User   *User
}

func (u *User) TableName() string {
	return "users"
}
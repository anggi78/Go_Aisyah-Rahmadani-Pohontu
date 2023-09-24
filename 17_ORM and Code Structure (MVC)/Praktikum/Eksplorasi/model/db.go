package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
    Name     string `json:"name" form:"name"`
    Email    string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
    Blogs    []Blog //`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"blogs"`
}
type Blog struct {
	gorm.Model
    Title   string `json:"title" form:"title"`
    Content string `json:"content" form:"content"`
    UserID  uint    `json:"user_id"`
    User    *User
}
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Book struct {
	gorm.Model
	Judul    string `json:"judul" form:"judul"`
	Penulis string `json:"penulis" form:"penulis"`
	Penerbit string `json:"penerbit" form:"penerbit"`
}
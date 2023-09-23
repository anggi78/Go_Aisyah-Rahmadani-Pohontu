package main

import (
	"Soal2/config"
	"Soal2/route"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()

	e := route.New()
	e.Logger.Fatal(e.Start(":8000"))
}
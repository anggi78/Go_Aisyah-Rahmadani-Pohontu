package main

import (
	"project/config"
	m "project/middleware"
	"project/route"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	config.InitDB()

	e := route.New()
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}

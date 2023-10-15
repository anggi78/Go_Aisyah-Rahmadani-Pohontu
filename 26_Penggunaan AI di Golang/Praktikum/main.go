package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"task/handler"
	"task/usecase"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	laptopUsecase := usecase.NewLaptopUsecase()
	laptopHandler := handler.NewLaptopHandler(laptopUsecase)
	e.POST("/recommended", laptopHandler.RecommendLaptop)
	port := ":8000"
	e.Start(port)
}
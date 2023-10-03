package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/repositories"
	"belajar-go-echo/router"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	userRepository := repositories.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)

	router.NewRouter(e, userController)

	e.Start(":8080")
}

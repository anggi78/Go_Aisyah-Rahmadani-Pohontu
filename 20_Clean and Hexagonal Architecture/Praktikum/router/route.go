package router

import (
	"belajar-go-echo/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, uc *controller.UserController) {
	AppJWT := e.Group("/users")
	AppJWT.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	AppJWT.GET("", uc.GetAllUsers)
	e.POST("/users", uc.CreateUser)
}
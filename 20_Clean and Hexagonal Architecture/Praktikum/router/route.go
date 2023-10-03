package router

import (
	"belajar-go-echo/constants"
	"belajar-go-echo/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, uc *controller.UserController) {
	AppJWT := e.Group("/users")
	AppJWT.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	AppJWT.GET("", uc.GetAllUsers)
	e.POST("/users", uc.CreateUser)
}

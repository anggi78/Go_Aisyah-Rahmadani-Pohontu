package route

import (
	"project/controller"
	"project/middleware"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/users/login", controller.LoginUserController)
	e.POST("/users", controller.CreateUserController)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWTMiddleware())

	eJwt.GET("/users", controller.GetUsersController)
	eJwt.GET("/users/:id", controller.GetUserController)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)

	eJwt.GET("/books", controller.GetBooksController)
	eJwt.GET("/books/:id", controller.GetBookController)
	eJwt.POST("/books", controller.CreateBookController)
	eJwt.DELETE("/books/:id", controller.DeleteBookController)
	eJwt.PUT("/books/:id", controller.UpdateBookController)
	return e
}

package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserController(userUseCase *usecase.UserUseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (uc *UserController) GetAllUsers(c echo.Context) error {
	users, err := uc.UserUseCase.GetAllUsers()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (uc *UserController) CreateUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	if err := uc.UserUseCase.CreateUser(user); err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": user,
	})
}

func (uc *UserController) LoginUser(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	token, err := uc.UserUseCase.LoginUser(user)
	if err != nil {
		return c.JSON(401, echo.Map{
			"error": "Authentication failed",
		})
	}

	return c.JSON(200, echo.Map{
		"token": token,
	})
}

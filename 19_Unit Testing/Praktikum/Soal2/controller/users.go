package controller

import (
	"net/http"
	"project/config"
	"project/middleware"

	"project/model"

	"github.com/labstack/echo/v4"
)

// users
func GetUsersController(c echo.Context) error {
	var users []model.User
	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var UserResponse = []model.UserResponse{}
	for _, v := range users {
		user := model.UserResponse{
			ID: v.ID,
			Name : v.Name,
			Email : v.Email,
		}
		UserResponse = append(UserResponse, user)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   UserResponse,
	})
}

func GetUserController(c echo.Context) error {
	userID := c.Param("id")
	var user model.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	UserResponse  := model.UserResponse{}
	UserResponse.ID = user.ID
	UserResponse.Name = user.Name
	UserResponse.Email = user.Email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by ID",
		"user":    UserResponse,
	})
}

func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	UserResponse  := model.UserResponse{}
	UserResponse.ID = user.ID
	UserResponse.Name = user.Name
	UserResponse.Email = user.Email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    UserResponse,
	})
}

func DeleteUserController(c echo.Context) error {
	userID := c.Param("id")
	var user model.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

func UpdateUserController(c echo.Context) error {
	userID := c.Param("id")
	var user model.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	updateUser := new(model.User)
	if err := c.Bind(updateUser); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user.Name = updateUser.Name
	user.Email = updateUser.Email
	user.Password = updateUser.Password

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	UserResponse  := model.UserResponse{}
	UserResponse.ID = user.ID
	UserResponse.Name = user.Name
	UserResponse.Email = user.Email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    UserResponse,
	})
}

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Login failed",
		})
	}

	token, err := middleware.CreateToken(int(user.ID))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	UserResponse  := model.UserResponse{}
	UserResponse.ID = user.ID
	UserResponse.Name = user.Name
	UserResponse.Email = user.Email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    UserResponse,
		"token":   token,
	})
}

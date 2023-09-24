package controller

import (
	"blogs/config"
	"blogs/model"
	"net/http"

	"github.com/labstack/echo"
)

// users
func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

func GetUserController(c echo.Context) error {
	userID := c.Param("id")
	var user model.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by ID",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
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
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    user,
	})
}

// blogs
func GetBlogsController(c echo.Context) error {
	var blogs []model.Blog
	if err := config.DB.Preload("User").Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"blogs": blogs,
	})
}

func GetBlogController(c echo.Context) error {
	blogID := c.Param("id")
	var blog model.Blog
	if err := config.DB.Preload("User").First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"blog": blog,
	})
}

func CreateBlogController(c echo.Context) error {
	blog := model.Blog{}
	if err := c.Bind(&blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user model.User
	if err := config.DB.Where("id = ?", blog.UserID).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong UserID")
	}

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create blog",
	})
}

func DeleteBlogController(c echo.Context) error {
	blogID := c.Param("id")
	var blog model.Blog
	if err := config.DB.First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}
	if err := config.DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog",
	})
}

func UpdateBlogController(c echo.Context) error {
	blogID := c.Param("id")
	var blog model.Blog
	if err := config.DB.First(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}

	updateBlog := new(model.Blog)
	if err := c.Bind(updateBlog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	blog.Title = updateBlog.Title
	blog.Content = updateBlog.Content
	blog.UserID = updateBlog.UserID

	if err := config.DB.Save(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := config.DB.Preload("User").Find(&blog, blogID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"blog": blog,
	})
}
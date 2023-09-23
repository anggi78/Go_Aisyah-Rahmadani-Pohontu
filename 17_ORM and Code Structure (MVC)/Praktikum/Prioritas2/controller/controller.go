package controller

import (
	"books/config"
	"books/model"
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

// books
func GetBooksController(c echo.Context) error {
	var books []model.Book
	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"books":   books,
	})
}

func GetBookController(c echo.Context) error {
	bookID := c.Param("id")
	var book model.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get book by ID",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {
	book := model.Book{}
	c.Bind(&book)
	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	bookID := c.Param("id")
	var book model.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}
	if err := config.DB.Delete(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete book",
	})
}

func UpdateBookController(c echo.Context) error {
	bookID := c.Param("id")
	var book model.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	}

	updateBook := new(model.Book)
	if err := c.Bind(updateBook); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	book.Judul = updateBook.Judul
	book.Penulis = updateBook.Penerbit
	book.Penerbit = updateBook.Penerbit

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book",
		"book":    book,
	})
}

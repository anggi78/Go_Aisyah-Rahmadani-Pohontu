package controller

import (
	"net/http"
	"project/config"

	"project/model"

	"github.com/labstack/echo/v4"
)

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
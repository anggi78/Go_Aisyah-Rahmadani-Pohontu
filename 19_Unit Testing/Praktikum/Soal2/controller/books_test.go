package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/model"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InsertDataBooks() error {
	user := model.Book{
		Judul:     "Book Baru",
		Penulis:    "Penulis Buku",
		Penerbit: "Penerbit Buku",
	}
	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetBooksController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetBooksController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")
}

func TestGetBookController(t *testing.T) {
    e := echo.New()

	userID := "1"
	InsertDataBooks()

	req := httptest.NewRequest(http.MethodGet, "/books/"+userID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetBookController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")
}


func TestCreateBookController(t *testing.T) {
	e := echo.New()

	reqBody := `{
        "Judul": "semangat ngoding",
        "Penulis": "anggi",
        "Penerbit": "aisyah"
    }`

	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateBookController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")

	var response struct {
		Message string     `json:"message"`
		Book    model.Book `json:"book"`
	}
	err := json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "success create new book", response.Message)
	assert.Equal(t, "semangat ngoding", response.Book.Judul)
	assert.Equal(t, "anggi", response.Book.Penulis)
	assert.Equal(t, "aisyah", response.Book.Penerbit)
}

func TestDeleteBookController(t *testing.T) {
    e := echo.New()

    InsertDataBooks()

    bookID := "1"

    req := httptest.NewRequest(http.MethodDelete, "/books/"+bookID, nil)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    err := DeleteBookController(c)

    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Contains(t, rec.Body.String(), "success delete book")
}

func TestUpdateBookController(t *testing.T) {
	e := echo.New()

	userID := "1"
	InsertDataBooks()

	reqBody := `{
		"Judul": "anggiUpdate",
		"Penulis": "new",
		"Penerbit": "new"
	}` 

	req := httptest.NewRequest(http.MethodPut, "/books/"+userID, strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	UpdateBookController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")
}
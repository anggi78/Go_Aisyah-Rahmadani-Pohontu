package controller

import (
	"net/http"
	"net/http/httptest"
	"project/config"
	"project/model"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitEchoTestAPI() *echo.Echo {
	config.ConnectDBTest()
	e := echo.New()
	return e
}

func InsertDataUserForGetUsers() error {
	user := model.User{
		Name:     "aisyah",
		Password: "123",
		Email:    "aisyah@gmail.com",
	}
	var err error
	if err = config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func TestGetUsersController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetUsersController(c)

	InsertDataUserForGetUsers()


	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")
}

func TestGetUserController(t *testing.T) {
	e := echo.New()

	userID := "1"
	InsertDataUserForGetUsers()

	req := httptest.NewRequest(http.MethodGet, "/users/"+userID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	GetUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")
}

func TestCreateUserController(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateUserController(c)

	InsertDataUserForGetUsers()

	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	CreateUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")
}

func TestDeleteUserController(t *testing.T) {
	e := echo.New()

	userID := "1"
	InsertDataUserForGetUsers()

	req := httptest.NewRequest(http.MethodDelete, "/users/"+userID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	DeleteUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")
}

func TestUpdateUserController(t *testing.T) {
	e := echo.New()

	userID := "1"
	InsertDataUserForGetUsers()

	reqBody := `{
		"name": "anggiUpdate",
		"email": "anggiUpdate@gmail.com",
		"password": "new"
	}` 

	req := httptest.NewRequest(http.MethodPut, "/users/"+userID, strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	UpdateUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")

	req = httptest.NewRequest(http.MethodPut, "/users/99", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	UpdateUserController(c)
}

func TestLoginUserController(t *testing.T) {
	e := echo.New()

	InsertDataUserForGetUsers()

	reqBody := `{
		"email": "aisyah@gmail.com",
		"password": "123"
	}`

	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	LoginUserController(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "success")

	invalidReqBody := `{
        "email": "invalid_email@gmail.com",
        "password": "wrong password"
    }`

    req = httptest.NewRequest(http.MethodPost, "/users/login", strings.NewReader(invalidReqBody))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec = httptest.NewRecorder()
    c = e.NewContext(req, rec)

    LoginUserController(c)

    assert.Equal(t, http.StatusBadRequest, rec.Code)
    assert.Contains(t, rec.Body.String(), "error")
}
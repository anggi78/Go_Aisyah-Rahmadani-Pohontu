package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	//"belajar-go-echo/helpers"
	"belajar-go-echo/model"
	"belajar-go-echo/repositories"
	"belajar-go-echo/usecase"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUserController_GetAllUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUserRepository := &repositories.MockUserRepository{} 
	mockUserUseCase := repositories.NewMockUserUseCase(mockUserRepository)
	userUseCase := usecase.NewUserUseCase(mockUserUseCase)

	userController := NewUserController(userUseCase)

	mockUsers := []model.User{
        {Name: "anggi", Email: "anggi@gmail.com", Password: "123"},
        {Name: "aisyah", Email: "aisyah@gmail.com", Password: "222"},
    }
    mockUserRepository.On("GetAllUsers").Return(mockUsers, nil)
	err := userController.GetAllUsers(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUserController_CreateUser(t *testing.T) {
    e := echo.New()
    reqBody := `{"Name": "anggi", "Email": "anggi@gmail.com", "Password": "123"}`
    req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(reqBody))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    mockUserRepository := &repositories.MockUserRepository{}
    mockUserUseCase := usecase.NewUserUseCase(mockUserRepository)
    userController := NewUserController(mockUserUseCase)

    user := &model.User{Name: "anggi", Email: "anggi@gmail.com", Password: "123"}
    mockUserRepository.On("CreateUser", user).Return(nil)

    err := userController.CreateUser(c)

    assert.NoError(t, err)
    assert.Equal(t, http.StatusOK, rec.Code)

    mockUserRepository.AssertCalled(t, "CreateUser", user)
}

// func TestUserController_LoginUser(t *testing.T) {
//     e := echo.New()
//     reqBody := `{"Name": "anggi", "Password": "123"}`
//     req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
//     req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//     rec := httptest.NewRecorder()
//     c := e.NewContext(req, rec)

//     mockUserRepository := &repositories.MockUserRepository{}
//     mockUserUseCase := usecase.NewUserUseCase(mockUserRepository)
//     userController := NewUserController(mockUserUseCase)

//     user := &model.User{Name: "anggi", Password: "123"}
//     token := "anggikiyowocantikmanisimutwawdaebak"
//     mockUserRepository.On("LoginUser", user).Return(nil)
//     mockUserRepository.On(helpers.GenerateToken(user.ID)).Return(token, nil)

//     err := userController.LoginUser(c)

//     assert.NoError(t, err)
//     assert.Equal(t, http.StatusOK, rec.Code)

//     mockUserRepository.AssertCalled(t, "LoginUser", user)
//    mockUserRepository.AssertCalled(t, "GenerateToken", user.ID)
// }
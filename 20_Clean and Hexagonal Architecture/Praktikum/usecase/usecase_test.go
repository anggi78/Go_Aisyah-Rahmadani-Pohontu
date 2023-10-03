package usecase

import (
	"belajar-go-echo/model"
	"belajar-go-echo/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserUseCase_CreateUser(t *testing.T) {
	userRepository := new(repositories.MockUserRepository)
	userUseCase := NewUserUseCase(userRepository)

	user := &model.User{Name: "anggi kiyowo", Password: "password123"}

	userRepository.On("CreateUser", user).Return(nil)

	err := userUseCase.CreateUser(user)
	assert.NoError(t, err)
}

func TestUserUseCase_GetAllUsers(t *testing.T) {
	userRepository := new(repositories.MockUserRepository)
	userUseCase := NewUserUseCase(userRepository)

	usersToReturn := []model.User{
		{Name: "anggi cantik", Password: "password1"},
		{Name: "anggi manis", Password: "password2"},
	}
	userRepository.On("GetAllUsers").Return(usersToReturn, nil)

	users, err := userUseCase.GetAllUsers()
	assert.NoError(t, err)
	assert.Equal(t, usersToReturn, users)
}

// func TestUserUseCase_LoginUser(t *testing.T) {
// 	userRepository := new(repositories.MockUserRepository)
// 	userUseCase := NewUserUseCase(userRepository)

// 	user := &model.User{Name: "anggi kiyowo", Password: "password123"}
// 	token := "anggikiyowocantikmanisimutwawdaebak"

// 	userRepository.On("LoginUser", user).Return(nil)
// 	userRepository.On("helpers.GenerateToken", user.ID).Return(token, nil)

// 	returnedToken, err := userUseCase.LoginUser(user)
// 	assert.NoError(t, err)
// 	assert.Equal(t, token, returnedToken)
// }

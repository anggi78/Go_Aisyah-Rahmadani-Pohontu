package usecase

import (
	"belajar-go-echo/helpers"
	"belajar-go-echo/model"
	"belajar-go-echo/repositories"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func NewUserUseCase(userRepository repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *UserUseCase) CreateUser(user *model.User) error {
	err := uc.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (uc *UserUseCase) GetAllUsers() ([]model.User, error) {
	users, err := uc.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uc *UserUseCase) LoginUser(user *model.User) (string, error) {
	err := uc.UserRepository.LoginUser(user)
	if err != nil {
		return "", err
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
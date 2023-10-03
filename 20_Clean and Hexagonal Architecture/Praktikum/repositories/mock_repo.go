package repositories

import (
    "belajar-go-echo/model"
    "github.com/stretchr/testify/mock"
)

//User Repository
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) CreateUser(user *model.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUserRepository) GetAllUsers() ([]model.User, error) {
    args := m.Called()
    return args.Get(0).([]model.User), args.Error(1)
}

func (m *MockUserRepository) LoginUser(user *model.User) error {
    args := m.Called(user)
    return args.Error(0)
}

//User UseCase
type MockUserUseCase struct {
    UserRepository *MockUserRepository
}

func NewMockUserUseCase(userRepository *MockUserRepository) *MockUserUseCase {
    return &MockUserUseCase{
        UserRepository: userRepository,
    }
}

func (mu *MockUserUseCase) CreateUser(user *model.User) error {
    return mu.UserRepository.CreateUser(user)
}

func (mu *MockUserUseCase) GetAllUsers() ([]model.User, error) {
    return mu.UserRepository.GetAllUsers()
}

func (mu *MockUserUseCase) LoginUser(user *model.User) error {
    return mu.UserRepository.LoginUser(user)
}
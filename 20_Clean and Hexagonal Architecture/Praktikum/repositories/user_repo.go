package repositories

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]model.User, error)
	LoginUser(user *model.User) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(user *model.User) error {
	result := ur.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ur *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := ur.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (ur *userRepository) LoginUser(user *model.User) error {
	result := ur.DB.Where(&model.User{Name: user.Name, Password: user.Password}).First(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

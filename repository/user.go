package repository

import (
	"project_sdu/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) (model.User, error)
	GetUserByID(id int) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}
func (u *userRepository) Add(user model.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil 
}

func (u *userRepository) CheckAvail(user model.User) (model.User, error) {
	var userExist model.User

	if err := u.db.Where("email = ?", user.Email).First(&userExist).Error; err != nil {
		return model.User{}, err
	}

	return userExist, nil 
}

func (u *userRepository) GetUserByID(id int) (model.User, error) {
	var user model.User

	if err := u.db.Where("id = ?", id).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil 
}

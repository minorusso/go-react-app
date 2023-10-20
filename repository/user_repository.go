package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

// インターフェースはメソッドの羅列
type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error //返り値の型
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

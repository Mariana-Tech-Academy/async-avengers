package repository

import (
	//"invoiceSys/db" dependency injection using r*

	"invoiceSys/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateCakeItem(user *models.User) error
	UpdateUser(user *models.User) error
}

type UserRepo struct {
	DB *gorm.DB
}

// Handler layer-------->services layer-------->repository layer(db methods)

func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (r *UserRepo) CreateCakeItem(user *models.User) error {
	err := r.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UpdateUser(user *models.User) error { // 1.2 update business profile
	return r.DB.Save(user).Error
}


package repository

import (
	"invoiceSys/db"
	"invoiceSys/models"
)

type BusinessRepository interface {
CreateBusiness(business *models.Business) error
GetBusinessByUserID(userID uint) (*models.Business, error)
UpdateBusiness(business *models.Business) error
}

type BusinessRepo struct{}

func (r *BusinessRepo) CreateBusiness(business *models.Business) error {
	return db.DB.Create(business).Error // // We use r.DB (the local struct field) instead of a global db.DB
}

func (r *BusinessRepo) GetBusinessByUserID(userID uint) (*models.Business, error) {
	var business models.Business
	err := db.DB.Where("user_id = ?", userID).First(&business).Error
	if err != nil {
		return nil, err
	}
	return &business, nil
}

func (r *BusinessRepo) UpdateBusiness(business *models.Business) error {
	return db.DB.Save(business).Error
}

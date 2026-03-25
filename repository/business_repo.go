package repository

import (
	"invoiceSys/db"
	//"invoiceSys/db"
	"invoiceSys/models"

	"gorm.io/gorm"
)

// Defining the database operations for business profiles
type BusinessRepository interface {
	CreateBusiness(business *models.Business) error            // US 1.1 saves a new business profile to db
	GetBusinessByUserID(userID uint) (*models.Business, error) // US 1.1 retrieves business profile so it can appear on invoices
	UpdateBusiness(business *models.Business) error            // US 1.2 updates existing business profile in the db
}

type BusinessRepo struct {
	DB *gorm.DB
}

// US 1.1 Create business profile - saves the new business to db
func (r *BusinessRepo) CreateBusiness(business *models.Business) error {
	return db.DB.Create(business).Error // // We use r.DB (the local struct field) instead of a global db.DB
}

// US 1.1 Business details appear on invoices - gets business profile by user ID
func (r *BusinessRepo) GetBusinessByUserID(userID uint) (*models.Business, error) {
	var business models.Business
	err := db.DB.Where("user_id = ?", userID).First(&business).Error
	if err != nil {
		return nil, err
	}
	return &business, nil
}

// US 1.2 Edit business info - saves updated business details to db
func (r *BusinessRepo) UpdateBusiness(business *models.Business) error {
	return db.DB.Save(business).Error
}

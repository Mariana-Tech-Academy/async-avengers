package repository

import (
	"invoiceSys/db"
	"invoiceSys/models"
)

// Defining the database operations for products
type ProductRepository interface {
	CreateProduct(product *models.Product) error                  // US 3.1 - saves a new product to the database
	GetProductByID(productID uint) (*models.Product, error)       // US 3.1 - retrieves a product by ID
	GetProductsByUserID(userID uint) ([]models.Product, error)    // US 3.1 - retrieves all products for a business owner
	UpdateProduct(product *models.Product) error                  // US 3.2 - updates existing product in the database
}

type ProductRepo struct{}

// US 3.1 - Create product (saves new product to the database)
func (r *ProductRepo) CreateProduct(product *models.Product) error {
	err := db.DB.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

// US 3.1 - Get product by ID
func (r *ProductRepo) GetProductByID(productID uint) (*models.Product, error) {
	var product models.Product
	err := db.DB.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return &models.Product{}, err
	}
	return &product, nil
}

// US 3.1 - Get all products
func (r *ProductRepo) GetProductsByUserID(userID uint) ([]models.Product, error) {
	var products []models.Product
	err := db.DB.Where("user_id = ?", userID).Find(&products).Error
	if err != nil {
		return []models.Product{}, err
	}
	return products, nil
}

// US 3.2 - Edit product
func (r *ProductRepo) UpdateProduct(product *models.Product) error {
	err := db.DB.Save(&product).Error
	if err != nil {
		return err
	}
	return nil
}
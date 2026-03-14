package repository

import (
    "invoiceSys/models"
    "gorm.io/gorm"
)

// ProductRepository defines the contract for product database operations
type ProductRepository interface {
    CreateProduct(product *models.Product) error
    UpdateProduct(product *models.Product) error
    GetProductByID(id uint) (*models.Product, error)
    GetProductsByBusinessID(businessID uint) ([]models.Product, error)
}

type ProductRepo struct {
    DB *gorm.DB
}

// CreateProduct persists a new product to the database (US 3.1)
func (r *ProductRepo) CreateProduct(product *models.Product) error {
    return r.DB.Create(product).Error
}

// UpdateProduct updates an existing product's details (US 3.2)
// GORM's Save method updates all fields of the provided model
func (r *ProductRepo) UpdateProduct(product *models.Product) error {
    return r.DB.Save(product).Error
}

// GetProductByID fetches a single product by its primary key
func (r *ProductRepo) GetProductByID(id uint) (*models.Product, error) {
    var product models.Product
    err := r.DB.First(&product, id).Error
    if err != nil {
        return nil, err
    }
    return &product, nil
}

// GetProductsByBusinessID retrieves all products belonging to a specific business
func (r *ProductRepo) GetProductsByBusinessID(businessID uint) ([]models.Product, error) {
    var products []models.Product
    err := r.DB.Where("business_id = ?", businessID).Find(&products).Error
    return products, err
}
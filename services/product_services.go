package services

import (
    "errors"
    "invoiceSys/models"
    "invoiceSys/repository"
)

type ProductService struct {
    Repo repository.ProductRepository
}

// US 3.1 – Create Product or Service
func (s *ProductService) CreateProduct(product *models.Product) error {
    // Business Logic: Ensure price isn't negative
    if product.Price < 0 {
        return errors.New("price cannot be less than zero")
    }

    if product.Name == "" {
        return errors.New("product name is required")
    }

    return s.Repo.CreateProduct(product)
}

// US 3.2 – Edit Product or Service
func (s *ProductService) UpdateProduct(productID uint, updatedData *models.Product) error {
    // 1. Fetch existing product to verify it exists
    existing, err := s.Repo.GetProductByID(productID)
    if err != nil {
        return errors.New("product not found")
    }

    // 2. Modify only allowed fields based on Acceptance Criteria
    // This ensures historical data in other tables isn't corrupted 
    // and only future calls to this product get the new info.
    existing.Price = updatedData.Price
    existing.Description = updatedData.Description
    
    // If the name is also allowed to be modified:
    if updatedData.Name != "" {
        existing.Name = updatedData.Name
    }

    // 3. Save the updated version
    return s.Repo.UpdateProduct(existing)
}

// Helper to list products for a business
func (s *ProductService) GetBusinessProducts(businessID uint) ([]models.Product, error) {
    return s.Repo.GetProductsByBusinessID(businessID)
}
package services

import (
	"errors"
	"invoiceSys/models"
	"invoiceSys/repository"
)

type ProductService struct {
	Repo repository.ProductRepository
}

// Handles the logic for creating a new product
func (s *ProductService) CreateProduct(product *models.Product) error {
	// save the new product to the database
	err := s.Repo.CreateProduct(product)
	if err != nil {
		return err
	}
	return nil
}

// Get product by ID
func (s *ProductService) GetProductByID(productID uint) (*models.Product, error) {
	product, err := s.Repo.GetProductByID(productID)
	if err != nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

// Get all products
func (s *ProductService) GetProductsByUserID(userID uint) ([]models.Product, error) {
	products, err := s.Repo.GetProductsByUserID(userID)
	if err != nil {
		return nil, errors.New("no products found")
	}
	return products, nil
}

// Updating - Handles the logic for updating product details
func (s *ProductService) UpdateProduct(productID uint, updated *models.Product) error {
	// fetch the existing product before updating
	existing, err := s.Repo.GetProductByID(productID)
	if err != nil {
		return errors.New("product not found")
	}

	// Updates price and description
	existing.Name = updated.Name
	existing.Price = updated.Price
	existing.Description = updated.Description

	// Saves the updated product
	err = s.Repo.UpdateProduct(existing)
	if err != nil {
		return err
	}
	return nil
}
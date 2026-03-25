package models

import "gorm.io/gorm"

// Product or service that can be added to an invoice
type Product struct {
	gorm.Model

	// US 3.1 - Create Product or Service
	UserID      uint    `json:"user_id"`     // links the product to a specific business owner
	Name        string  `json:"name"`        
	Price       float64 `json:"price"`       
	Description string  `json:"description"` 
}
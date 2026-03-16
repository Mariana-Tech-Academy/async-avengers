package models

import (
	"gorm.io/gorm"
)


type Product struct { //generic cake for modification
	gorm.Model
	Name        string  `json:"name"`        // e.g., "Signature Chocolate Cake"
	Price       float64 `json:"price"`       // e.g., 60.00
	Description string  `json:"description"` // e.g., "Triple layer with ganache"
	UserID      uint    `json:"user_id"`     // Who created this product
}

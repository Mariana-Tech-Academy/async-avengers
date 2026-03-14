package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ID          uint    `json:"id"`
	BusinessID  uint    `json:"business_id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

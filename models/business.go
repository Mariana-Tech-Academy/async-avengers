package models

import (

	"gorm.io/gorm"
)


// Creating business profile
type Business struct {
	gorm.Model

	UserID  uint     `json:"user_id"`
	Name    string   `json:"name"`
	Address string   `json:"address"`
	Phone   string   `json:"phone"`
	Email   string   `json:"email"`

// Business logo stored as a file path or URL
Logo string `json:"Logo"`

// Adding Tax information
VATNumber string  `json:"vat_number"`
TaxRate   float64 `json:"tax_rate"`  

}

package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string  `json:"username"`
	Password        string  `json:"password"`
	Buisnessname    string  `json:"buisnessname"`
	Buisnessaddress string  `json:"buisnessaddress"`
	Phone           string  `json:"phone"`
	Email           string  `json:"email" gorm:"uniqueIndex"`
	TaxID           string  `json:"tax_id"`
	TaxRate         float64 `json:"tax_rate"`
}

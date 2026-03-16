
package models

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Invoices []Invoice `json:"invoices" gorm:"foreignKey:ClientID"` // Added for US 2.3
}

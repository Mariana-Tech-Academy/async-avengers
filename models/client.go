package models

import (
	"gorm.io/gorm"
)

// Client model - a customer that the business invoices
type Client struct {
	gorm.Model

	// US 2.1 - Add Client
	UserID   uint      `json:"user_id"`  // links the client to a specific business owner
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
}
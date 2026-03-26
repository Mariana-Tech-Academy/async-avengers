package models

import "gorm.io/gorm"

// Client model - a customer that the business invoices
type Client struct {
	gorm.Model

	// US 2.1 - Add Client
	UserID  uint   `json:"-"`       // links the client to a specific business owner
	Name    string `json:"name"`    // client name
	Email   string `json:"email"`   // client email
	Phone   string `json:"phone"`    // client phone number
	Address string `json:"address"` // client billing address
}

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

type Invoice struct {
	gorm.Model
	UserID        uint      `json:"user_id"`
	ClientID      uint      `json:"client_id"` // Link to Client
	Subtotal      float64   `json:"subtotal"`
	ServiceCharge float64   `json:"service_charge"`
	TaxRate       float64   `json:"tax_rate"`
	TaxAmount     float64   `json:"tax_amount"`
	TotalAmount   float64   `json:"total_amount"`
	Status        string    `json:"status"` // e.g., "Pending", "Paid" (Required for US 2.3)
	Items         []CakeItem `json:"items" gorm:"foreignKey:InvoiceID"`
}

type Product struct { //generic cake for modification
	gorm.Model
	Name        string  `json:"name"`        // e.g., "Signature Chocolate Cake"
	Price       float64 `json:"price"`       // e.g., 60.00
	Description string  `json:"description"` // e.g., "Triple layer with ganache"
	UserID      uint    `json:"user_id"`     // Who created this product
}

type Client struct {
	gorm.Model
	Name     string    `json:"name"`
	Email    string    `json:"email" gorm:"uniqueIndex"`
	Phone    string    `json:"phone"`
	Address  string    `json:"address"`
	Invoices []Invoice `json:"invoices" gorm:"foreignKey:ClientID"` // Added for US 2.3
}

type CakeItem struct {
	gorm.Model
	InvoiceID uint `json:"invoice_id"`

	Size   string `json:"size"`   // in inches
	Flavor string `json:"flavor"` // roughly 3/4
	Filler string `json:"filler"` // 3/4 flavours

	Quantity      int     `json:"quantity"`
	UnitPrice     float64 `json:"unit_price"`
	ServiceCharge float64 `json:"service_charge"`
	TaxRate       float64 `json:"tax_rate"`
	Total         float64 `json:"total"`

}

func (c *CakeItem) CalculateTotal() {
	subtotal := (float64(c.Quantity) * c.UnitPrice) + c.ServiceCharge
	tax := subtotal * (c.TaxRate / 100)
	c.Total = subtotal + tax
}


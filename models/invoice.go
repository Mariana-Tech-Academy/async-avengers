
package models

import (
	"gorm.io/gorm"
)


type Invoice struct {
	gorm.Model
	UserID        uint       `json:"user_id"`
	ClientID      uint       `json:"client_id"` // Link to Client
	Subtotal      float64    `json:"subtotal"`
	ServiceCharge float64    `json:"service_charge"`
	TaxRate       float64    `json:"tax_rate"`
	TaxAmount     float64    `json:"tax_amount"`
	TotalAmount   float64    `json:"total_amount"`
	Status        string     `json:"status"` // e.g., "Pending", "Paid" (Required for US 2.3)
	Items         []CakeItem `json:"items" gorm:"foreignKey:InvoiceID"`
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
	TaxAmount     float64 `json:"tax_amount"`
	Total         float64 `json:"total"`
}

func (c *CakeItem) CalculateTotal() {
	const taxRate = 0.075 // 7.5% tax rate repr
    // Calculate subtotal
    subtotal := (float64(c.Quantity) * c.UnitPrice) + c.ServiceCharge
    
    // Store values back into the struct
    c.TaxRate = taxRate
    c.TaxAmount = subtotal * taxRate
    c.Total = subtotal + c.TaxAmount
}

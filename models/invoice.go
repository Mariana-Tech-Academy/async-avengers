package models

import "gorm.io/gorm"

// Invoice model -  a full invoice sent to a client
type Invoice struct {
	gorm.Model

	// US 4.1 - Create Invoice
	UserID        uint          `json:"-"`        // links invoice to business owner
	ClientID      uint          `json:"client_id"`      // client this invoice is for
	InvoiceNumber string        `json:"invoice_number"` // unique invoice number
	Status        string        `json:"status"`         // Draft, Sent, Paid, Overdue

	// US 4.1 - Automatic calculations
	Subtotal    float64 `json:"subtotal"`     // sum of all line items
	TaxRate     float64 `json:"tax_rate"`     // tax rate from business profile
	TaxAmount   float64 `json:"tax_amount"`   // calculated tax amount
	TotalAmount float64 `json:"total_amount"` // subtotal + tax

	// US 4.2 - Multiple line items
	Items []InvoiceItem `json:"items" gorm:"foreignKey:InvoiceID"`
}

// Invoice Item model -  a single line item on an invoice
type InvoiceItem struct {
	gorm.Model

	// US 4.2 - Add Multiple Line Items
	InvoiceID uint    `json:"invoice_id"` // links item to invoice
	ProductID uint    `json:"product_id"` // links item to product
	Name      string  `json:"name"`       // product name at time of invoice
	Quantity  int     `json:"quantity"`   // US 4.2: quantity of product
	UnitPrice float64 `json:"unit_price"` // US 4.2: price per unit at time of invoice
	Total     float64 `json:"total"`      // US 4.2: quantity * unit price
}
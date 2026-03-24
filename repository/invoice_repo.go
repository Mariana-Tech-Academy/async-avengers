package repository

import (
	"invoiceSys/db"
	"invoiceSys/models"
)

// Invoice repo defines the database operations for invoices
type InvoiceRepository interface {
	CreateInvoice(invoice *models.Invoice) error                        // US 4.1: saves a new invoice to the database
	GetInvoiceByID(invoiceID uint) (*models.Invoice, error)             // US 4.1: retrieves an invoice by ID
	GetInvoicesByUserID(userID uint) ([]models.Invoice, error)          // US 4.1: retrieves all invoices for a business owner
	GetInvoicesByClientID(clientID uint) ([]models.Invoice, error)      // US 2.3: retrieves all invoices for a client
	UpdateInvoice(invoice *models.Invoice) error                        // US 4.3: updates an existing invoice
	GetProductByID(productID uint) (*models.Product, error)
}

type InvoiceRepo struct{}

// Create Invoice - saves the new invoice to the database
func (r *InvoiceRepo) CreateInvoice(invoice *models.Invoice) error {
	err := db.DB.Create(&invoice).Error
	if err != nil {
		return err
	}
	return nil
}

// Get invoice by ID - fetches a single invoice with all its items
func (r *InvoiceRepo) GetInvoiceByID(invoiceID uint) (*models.Invoice, error) {
	var invoice models.Invoice
	err := db.DB.Preload("Items").Where("id = ?", invoiceID).First(&invoice).Error
	if err != nil {
		return &models.Invoice{}, err
	}
	return &invoice, nil
}

// Get all invoices for a user
func (r *InvoiceRepo) GetInvoicesByUserID(userID uint) ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := db.DB.Preload("Items").Where("user_id = ?", userID).Find(&invoices).Error
	if err != nil {
		return []models.Invoice{}, err
	}
	return invoices, nil
}

// View Client Invoice History - fetches all invoices for a client
func (r *InvoiceRepo) GetInvoicesByClientID(clientID uint) ([]models.Invoice, error) {
	var invoices []models.Invoice
	err := db.DB.Preload("Items").Where("client_id = ?", clientID).Find(&invoices).Error
	if err != nil {
		return []models.Invoice{}, err
	}
	return invoices, nil
}

// Save Invoice as Draft -  updates an existing invoice
func (r *InvoiceRepo) UpdateInvoice(invoice *models.Invoice) error {
	err := db.DB.Save(&invoice).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *InvoiceRepo) GetProductByID(productID uint) (*models.Product, error) {
	var product models.Product
	err := db.DB.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return &models.Product{}, err
	}
	return &product, nil
}
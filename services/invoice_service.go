package services

import (
	"errors"
	"fmt"
	"invoiceSys/db"
	"invoiceSys/models"
	"invoiceSys/repository"
	"time"
)

type InvoiceService struct {
	Repo         repository.InvoiceRepository
	BusinessRepo repository.BusinessRepository
	ProductRepo repository.ProductRepository
}

// Creates Invoice - handles the logic for creating a new invoice
func (s *InvoiceService) CreateInvoice(invoice *models.Invoice) error {

	// generates a unique invoice number e.g. INV-20260312-0001
	invoice.InvoiceNumber = fmt.Sprintf("INV-%s-%d",
		time.Now().Format("20060102"),
		time.Now().UnixNano()%10000,
	)

	// set default status to Draft
	invoice.Status = "Draft"

	// get tax rate from business profile
	business, err := s.BusinessRepo.GetBusinessByUserID(invoice.UserID)
	if err != nil {
		return errors.New("business profile not found")
	}
	invoice.TaxRate = business.TaxRate

	// calculate subtotal from all line items
	var subtotal float64
	for i := range invoice.Items {
		product, err := s.Repo.GetProductByID(invoice.Items[i].ProductID)
		if err != nil {
			return err
		}
		// calculate total for each line item
		invoice.Items[i].UnitPrice = product.Price
		invoice.Items[i].Total = float64(invoice.Items[i].Quantity) * product.Price
		subtotal += invoice.Items[i].Total
	}

	// calculate tax and total amount
	invoice.Subtotal = subtotal
	invoice.TaxAmount = subtotal * (invoice.TaxRate / 100)
	invoice.TotalAmount = subtotal + invoice.TaxAmount

	// save the invoice to the database
	err = s.Repo.CreateInvoice(invoice)
	if err != nil {
		return err
	}
	return nil
}

// Get invoice by ID - retrieves a single invoice
func (s *InvoiceService) GetInvoiceByID(invoiceID uint) (*models.Invoice, error) {
	invoice, err := s.Repo.GetInvoiceByID(invoiceID)
	if err != nil {
		return nil, errors.New("invoice not found")
	}
	return invoice, nil
}

// Get all invoices for a user
func (s *InvoiceService) GetInvoicesByUserID(userID uint) ([]models.Invoice, error) {
	invoices, err := s.Repo.GetInvoicesByUserID(userID)
	if err != nil {
		return nil, errors.New("no invoices found")
	}
	return invoices, nil
}

// View Client Invoice History - retrieves all invoices for a client
func (s *InvoiceService) GetInvoicesByClientID(clientID uint) ([]models.Invoice, error) {
	invoices, err := s.Repo.GetInvoicesByClientID(clientID)
	if err != nil {
		return nil, errors.New("no invoices found for this client")
	}
	return invoices, nil
}

// Save Invoice as Draft - updates an existing invoice
func (s *InvoiceService) UpdateInvoice(invoiceID uint, updated *models.Invoice) error {
	// fetch the existing invoice
	existing, err := s.Repo.GetInvoiceByID(invoiceID)
	if err != nil {
		return errors.New("invoice not found")
	}

	// only allow editing if invoice is still a Draft
	if existing.Status != "Draft" {
		return errors.New("only draft invoices can be edited")
	}

	// update invoice details
	existing.ClientID = updated.ClientID
	db.DB.Where("invoice_id = ?", existing.ID).Delete(&models.InvoiceItem{}) // delete old items first before adding new ones
	existing.Items = updated.Items

	// recalculate totals
	var subtotal float64
	for i := range existing.Items {
		existing.Items[i].Total = float64(existing.Items[i].Quantity) * existing.Items[i].UnitPrice
		subtotal += existing.Items[i].Total
	}
	existing.Subtotal = subtotal
	existing.TaxAmount = subtotal * (existing.TaxRate / 100)
	existing.TotalAmount = subtotal + existing.TaxAmount

	err = s.Repo.UpdateInvoice(existing)
	if err != nil {
		return err
	}
	return nil
}

// Mark Invoice as Paid - updates the invoice status to Paid
func (s *InvoiceService) UpdateInvoiceStatus(invoiceID uint, status string) error {
	existing, err := s.Repo.GetInvoiceByID(invoiceID)
	if err != nil {
		return errors.New("invoice not found")
	}
	existing.Status = status
	err = s.Repo.UpdateInvoice(existing)
	if err != nil {
		return err
	}
	return nil
}

// US 6.2 - View Invoice Status: returns the status of an invoice
func (s *InvoiceService) GetInvoiceStatus(invoiceID uint) (string, error) {
	invoice, err := s.Repo.GetInvoiceByID(invoiceID)
	if err != nil {
		return "", errors.New("invoice not found")
	}
	return invoice.Status, nil
}
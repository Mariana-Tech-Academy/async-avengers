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
	ProductRepo  repository.ProductRepository
}

// US 4.1 - Create Invoice: handles the logic for creating a new invoice
func (s *InvoiceService) CreateInvoice(invoice *models.Invoice) error {

	// US 4.1: generates a unique invoice number e.g. INV-20260312-0001
	invoice.InvoiceNumber = fmt.Sprintf("INV-%s-%d",
		time.Now().Format("20060102"),
		time.Now().UnixNano()%10000,
	)

	// US 4.3: set default status to Draft
	invoice.Status = "Draft"

	// US 4.1: get tax rate from business profile
	business, err := s.BusinessRepo.GetBusinessByUserID(invoice.UserID)
	if err != nil {
		return errors.New("business profile not found")
	}
	invoice.TaxRate = business.TaxRate

	// US 4.1 + US 4.2: calculate subtotal from all line items
	var subtotal float64
	for i := range invoice.Items {
		// automatically fetch product details if product_id is provided
		if invoice.Items[i].ProductID != 0 {
			product, err := s.ProductRepo.GetProductByID(invoice.Items[i].ProductID)
			if err == nil {
				// auto fill name and price from product
				invoice.Items[i].Name = product.Name
				invoice.Items[i].UnitPrice = product.Price
				// only use product description if none was provided
				if invoice.Items[i].Description == "" {
					invoice.Items[i].Description = product.Description
				}
			}
		}
		// US 4.2: calculate total for each line item
		invoice.Items[i].Total = float64(invoice.Items[i].Quantity) * invoice.Items[i].UnitPrice
		subtotal += invoice.Items[i].Total
	}

	// US 4.1: calculate tax and total amount
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

// US 4.1 - Get invoice by ID: retrieves a single invoice
func (s *InvoiceService) GetInvoiceByID(invoiceID uint) (*models.Invoice, error) {
	invoice, err := s.Repo.GetInvoiceByID(invoiceID)
	if err != nil {
		return nil, errors.New("invoice not found")
	}
	return invoice, nil
}

// US 4.1 - Get all invoices for a user
func (s *InvoiceService) GetInvoicesByUserID(userID uint) ([]models.Invoice, error) {
	invoices, err := s.Repo.GetInvoicesByUserID(userID)
	if err != nil {
		return nil, errors.New("no invoices found")
	}
	return invoices, nil
}

// US 2.3 - View Client Invoice History: retrieves all invoices for a client
func (s *InvoiceService) GetInvoicesByClientID(clientID uint) ([]models.Invoice, error) {
	invoices, err := s.Repo.GetInvoicesByClientID(clientID)
	if err != nil {
		return nil, errors.New("no invoices found for this client")
	}
	return invoices, nil
}

// US 4.3 - Save Invoice as Draft: updates an existing invoice
func (s *InvoiceService) UpdateInvoice(invoiceID uint, updated *models.Invoice) error {
	// fetch the existing invoice
	existing, err := s.Repo.GetInvoiceByID(invoiceID)
	if err != nil {
		return errors.New("invoice not found")
	}

	// US 4.3: only allow editing if invoice is still a Draft
	if existing.Status != "Draft" {
		return errors.New("only draft invoices can be edited")
	}

	// update invoice details
	existing.ClientID = updated.ClientID

	// delete old items first before adding new ones
	db.DB.Where("invoice_id = ?", existing.ID).Delete(&models.InvoiceItem{})
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

// US 6.1 - Mark Invoice as Paid: updates the invoice status
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
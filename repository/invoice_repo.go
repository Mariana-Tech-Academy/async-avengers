

package repository

import (
	//"invoiceSys/db" dependency injection using r*

	"invoiceSys/models"
	"gorm.io/gorm"
)


type InvoiceRepository interface {
	
	GetInvoicesByClientID(clientID uint) ([]models.Invoice, error) // 2.3 mx client
	CreateInvoice(invoice *models.Invoice) error                   //invoice items

}

type InvoiceRepo struct {
	DB *gorm.DB
}


func (r *InvoiceRepo) GetInvoicesByClientID(clientID uint) ([]models.Invoice, error) { //2.3 client invoice history to view
	var invoices []models.Invoice
	err := r.DB.Preload("Items").Where("client_id = ?", clientID).Find(&invoices).Error // get cake items assoc with specific invoicw
	return invoices, err
}

func (r *InvoiceRepo) CreateInvoice(invoice *models.Invoice) error { // 1.1 save cake items+ create invoice
	return r.DB.Create(invoice).Error
}

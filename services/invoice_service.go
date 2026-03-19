package services

import "invoiceSys/repository"

type InvoiceService struct {
	Repo *repository.InvoiceRepo
}

func (s *InvoiceService) SendInvoice(invoiceID uint, email string) error {

	PDFPath, err := GenerateInvoicePDF(invoiceID)
	if err != nil {
		return err
	}

	err = SendInvoiceEmail(email, PDFPath)
	if err != nil {
		return err
	}
	return nil
}
package services

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-pdf/fpdf"
	"invoiceSys/middleware"
	"invoiceSys/models"
	"invoiceSys/repository"
	"invoiceSys/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(req *models.User) error {

	// Check if user already exists
	_, err := s.Repo.GetUserByUsername(req.Username)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashedPass

	// Save user to DB
	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(req *models.User) (string, error) {

	// Check if user exists
	user, err := s.Repo.GetUserByUsername(req.Username)
	if err != nil {
		return "", err
	}

	// Compare password
	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil

}

func (s *UserService) DownloadInvoicePDF(invoiceID uint) ([]byte, error) {
	// fetch invoice details from the database
	invoice, err := s.Repo.GetInvoiceByID(invoiceID)
	if err != nil {
		return nil, err
	}
	// generate PDF using fpdf
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	// set font and add invoice details to the PDF
	pdf.SetFont("Arial", "B", 18)
	pdf.Cell(190, 10, "INVOICE")
	pdf.Ln(12)

	pdf.SetFont("Arial", "", 12)
	// add business and client details to the PDF
	pdf.Cell(100, 8, "Business: "+invoice.User.Buisnessname)
	pdf.Ln(6)

	pdf.Cell(100, 8, "Client: "+invoice.Client.Name)
	pdf.Ln(6)

	pdf.Cell(100, 8, "Email: "+invoice.Client.Email)
	pdf.Ln(10)
	//create table header for the invoice items
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(50, 8, "Item")
	pdf.Cell(20, 8, "Qty")
	pdf.Cell(30, 8, "Price")
	pdf.Cell(30, 8, "Tax")
	pdf.Cell(30, 8, "Total")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 11)
	// add invoice items to the PDF
	for _, item := range invoice.Items {

		name := item.Size + " " + item.Flavor

		pdf.Cell(50, 8, name)
		pdf.Cell(20, 8, fmt.Sprintf("%d", item.Quantity))
		pdf.Cell(30, 8, fmt.Sprintf("%.2f", item.UnitPrice))
		pdf.Cell(30, 8, fmt.Sprintf("%.2f", item.TaxAmount))
		pdf.Cell(30, 8, fmt.Sprintf("%.2f", item.Total))
		pdf.Ln(8)
	}

	pdf.Ln(5)
	// add invoice totals to the PDF
	pdf.Cell(120, 8, "")
	pdf.Cell(40, 8, "Subtotal:")
	pdf.Cell(30, 8, fmt.Sprintf("%.2f", invoice.Subtotal))
	pdf.Ln(6)

	pdf.Cell(120, 8, "")
	pdf.Cell(40, 8, "Tax:")
	pdf.Cell(30, 8, fmt.Sprintf("%.2f", invoice.TaxAmount))
	pdf.Ln(6)

	pdf.Cell(120, 8, "")
	pdf.Cell(40, 8, "Total:")
	pdf.Cell(30, 8, fmt.Sprintf("%.2f", invoice.TotalAmount))
	// convert the PDF to bytes (cannot download a go object it must receive a file data)
	var buf bytes.Buffer
	err = pdf.Output(&buf)
	if err != nil {
		return nil, err
	}
	// return the PDF bytes to the caller
	return buf.Bytes(), nil
}

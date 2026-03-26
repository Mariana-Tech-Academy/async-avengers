package services

import (
	"fmt"
	"invoiceSys/models"

	"github.com/jung-kurt/gofpdf"
)

// US 5.2 - Download Invoice as PDF: generates a PDF for a given invoice
func GenerateInvoicePDF(invoice *models.Invoice, business *models.Business, client *models.Client) (string, error) {

	// create a new A4 PDF document
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	// =====================
	// BUSINESS DETAILS SECTION (US 5.2 - PDF contains company info)
	// =====================
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(190, 10, business.Name)
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(190, 6, business.Address)
	pdf.Ln(6)
	pdf.Cell(190, 6, business.Phone)
	pdf.Ln(6)
	pdf.Cell(190, 6, business.Email)
	pdf.Ln(6)
	pdf.Cell(190, 6, fmt.Sprintf("VAT Number: %s", business.VATNumber))
	pdf.Ln(12)

	// =====================
	// INVOICE TITLE SECTION
	// =====================
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(190, 10, "INVOICE")
	pdf.Ln(8)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(190, 6, fmt.Sprintf("Invoice Number: %s", invoice.InvoiceNumber))
	pdf.Ln(6)
	pdf.Cell(190, 6, fmt.Sprintf("Status: %s", invoice.Status))
	pdf.Ln(12)

	// =====================
	// CLIENT DETAILS SECTION (US 5.2 - PDF contains client info)
	// =====================
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(190, 8, "Bill To:")
	pdf.Ln(7)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(190, 6, client.Name)
	pdf.Ln(6)
	pdf.Cell(190, 6, client.Email)
	pdf.Ln(6)
	pdf.Cell(190, 6, client.Address)
	pdf.Ln(12)

	// =====================
	// TABLE HEADER (US 5.2 - PDF contains item breakdown)
	// =====================
	pdf.SetFont("Arial", "B", 10)
	pdf.SetFillColor(200, 200, 200)
	pdf.CellFormat(50, 8, "Item", "1", 0, "", true, 0, "")
	pdf.CellFormat(70, 8, "Description", "1", 0, "", true, 0, "")
	pdf.CellFormat(15, 8, "Qty", "1", 0, "C", true, 0, "")
	pdf.CellFormat(27, 8, "Unit Price", "1", 0, "C", true, 0, "")
	pdf.CellFormat(28, 8, "Total", "1", 0, "C", true, 0, "")
	pdf.Ln(8)

	// =====================
	// LINE ITEMS (US 4.2 - multiple line items, US 5.2 - item breakdown)
	// =====================
	pdf.SetFont("Arial", "", 10)
	for _, item := range invoice.Items {
		x := pdf.GetX()
		y := pdf.GetY()

		// draw description cell first to calculate row height
		pdf.SetXY(x+50, y)
		pdf.MultiCell(70, 8, item.Description, "1", "", false)
		descHeight := pdf.GetY() - y

		// use description height as the row height for all cells
		rowHeight := descHeight

		// draw item name, quantity, unit price and total with same height
		pdf.SetXY(x, y)
		pdf.CellFormat(50, rowHeight, item.Name, "1", 0, "", false, 0, "")
		pdf.SetXY(x+120, y)
		pdf.CellFormat(15, rowHeight, fmt.Sprintf("%d", item.Quantity), "1", 0, "C", false, 0, "")
		pdf.CellFormat(27, rowHeight, fmt.Sprintf("GBP %.2f", item.UnitPrice), "1", 0, "C", false, 0, "")
		pdf.CellFormat(28, rowHeight, fmt.Sprintf("GBP %.2f", item.Total), "1", 0, "C", false, 0, "")
		pdf.SetXY(x, y+rowHeight)
	}
	pdf.Ln(4)

	// =====================
	// TOTALS SECTION (US 5.2 - PDF contains totals, US 1.3 - tax calculation)
	// =====================
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(150, 8, "Subtotal:")
	pdf.Cell(40, 8, fmt.Sprintf("GBP %.2f", invoice.Subtotal))
	pdf.Ln(8)

	pdf.Cell(150, 8, fmt.Sprintf("Tax (%.0f%%):", invoice.TaxRate))
	pdf.Cell(40, 8, fmt.Sprintf("GBP %.2f", invoice.TaxAmount))
	pdf.Ln(8)

	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(150, 8, "Total Amount:")
	pdf.Cell(40, 8, fmt.Sprintf("GBP %.2f", invoice.TotalAmount))
	pdf.Ln(12)

	// =====================
	// SAVE PDF FILE
	// =====================
	fileName := fmt.Sprintf("invoice_%s.pdf", invoice.InvoiceNumber)
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

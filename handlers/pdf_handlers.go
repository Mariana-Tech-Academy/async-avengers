package handlers

import (
	"invoiceSys/services"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type PDFHandler struct {
	InvoiceService  *services.InvoiceService
	BusinessService *services.BusinessService
	ClientService   *services.ClientService
}

// US 5.2 - Download Invoice as PDF: generates and downloads a PDF for a given invoice
func (h *PDFHandler) DownloadInvoicePDF(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceID, err := strconv.Atoi(vars["invoiceID"])
	if err != nil {
		http.Error(w, "invalid invoice ID", http.StatusBadRequest)
		return
	}

	// fetch the invoice
	invoice, err := h.InvoiceService.GetInvoiceByID(uint(invoiceID))
	if err != nil {
		http.Error(w, "invoice not found", http.StatusNotFound)
		return
	}

	// make sure invoice belongs to logged in user
	userID := r.Context().Value("user_id").(uint)
	if invoice.UserID != userID {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// fetch the business profile
	business, err := h.BusinessService.GetBusiness(invoice.UserID)
	if err != nil {
		http.Error(w, "business profile not found", http.StatusNotFound)
		return
	}

	// fetch the client
	client, err := h.ClientService.GetClientByID(invoice.ClientID)
	if err != nil {
		http.Error(w, "client not found", http.StatusNotFound)
		return
	}

	// generate the PDF
	fileName, err := services.GenerateInvoicePDF(invoice, business, client)
	if err != nil {
		http.Error(w, "failed to generate PDF", http.StatusInternalServerError)
		return
	}

	// US 5.2: send the PDF as a download
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)

	// read and send the PDF file
	pdfBytes, err := os.ReadFile(fileName)
	if err != nil {
		http.Error(w, "failed to read PDF", http.StatusInternalServerError)
		return
	}

	// delete the file after sending
	defer os.Remove(fileName)

	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}

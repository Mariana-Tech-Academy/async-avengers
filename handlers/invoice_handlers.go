package handlers

import (
	"encoding/json"
	"invoiceSys/models"
	"invoiceSys/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InvoiceHandler struct {
	Service *services.InvoiceService
}

// Create Invoice - receives the request and creates a new invoice
func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	// get user ID from token instead of request body
	userID := r.Context().Value("user_id").(uint)
	
	var invoice models.Invoice
	err := json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set user ID from token automatically
	invoice.UserID = userID

	// call service layer to create the invoice
	err = h.Service.CreateInvoice(&invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the created invoice
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

// Get invoice by ID - retrieves a single invoice
func (h *InvoiceHandler) GetInvoiceByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceID, err := strconv.Atoi(vars["invoiceID"])
	if err != nil {
		http.Error(w, "invalid invoice ID", http.StatusBadRequest)
		return
	}

	invoice, err := h.Service.GetInvoiceByID(uint(invoiceID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoice)
}

// Get all invoices for a user
func (h *InvoiceHandler) GetInvoicesByUserID(w http.ResponseWriter, r *http.Request) {
	// get user ID from token instead of URL
	userID := r.Context().Value("user_id").(uint)

	invoices, err := h.Service.GetInvoicesByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoices)
}

// View Client Invoice History - retrieves all invoices for a client
func (h *InvoiceHandler) GetInvoicesByClientID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, err := strconv.Atoi(vars["clientID"])
	if err != nil {
		http.Error(w, "invalid client ID", http.StatusBadRequest)
		return
	}

	invoices, err := h.Service.GetInvoicesByClientID(uint(clientID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(invoices)
}

// Save Invoice as Draft - updates an existing draft invoice
func (h *InvoiceHandler) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceID, err := strconv.Atoi(vars["invoiceID"])
	if err != nil {
		http.Error(w, "invalid invoice ID", http.StatusBadRequest)
		return
	}

	var updated models.Invoice
	err = json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateInvoice(uint(invoiceID), &updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

// Mark Invoice as Paid - updates the invoice status
func (h *InvoiceHandler) UpdateInvoiceStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceID, err := strconv.Atoi(vars["invoiceID"])
	if err != nil {
		http.Error(w, "invalid invoice ID", http.StatusBadRequest)
		return
	}

	// get the new status from the rb
	var body struct {
		Status string `json:"status"`
	}
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Service.UpdateInvoiceStatus(uint(invoiceID), body.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": body.Status})
}

// US 6.2 - View Invoice Status: returns the status of an invoice
func (h *InvoiceHandler) GetInvoiceStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	invoiceID, err := strconv.Atoi(vars["invoiceID"])
	if err != nil {
		http.Error(w, "invalid invoice ID", http.StatusBadRequest)
		return
	}

	status, err := h.Service.GetInvoiceStatus(uint(invoiceID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": status})
}
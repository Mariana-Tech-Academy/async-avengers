package handlers

import (
	"encoding/json"
	"invoiceSys/models"
	"invoiceSys/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InvoiceHandler struct{
    Service *services.InvoiceService
}

func (h *InvoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice //4.1 and 4.2 create invoice with multiple lines

	// 1. Collect client selection, items, and quantities from JSON
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateInvoice(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

func (h *InvoiceHandler) SendInvoice(w http.ResponseWriter, r *http.Request) {
	var req struct {
		InvoiceID uint   `json:"invoice_id"`
		Email     string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}
	if err := h.Service.SendInvoice(req.InvoiceID, req.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
        "message": "invoice sent!!",
    })
}

// US 4.3: Update/Save Draft
func (h *InvoiceHandler) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var updatedInvoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&updatedInvoice); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Service checks if status is 'Draft' before allowing certain edits
	if err := h.Service.UpdateInvoice(uint(id), &updatedInvoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedInvoice)
}

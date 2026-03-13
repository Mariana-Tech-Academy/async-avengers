package handlers

import (
	"encoding/json"
	"net/http"

	"invoiceSys/services"
)

type InvoiceRequest struct {
	Email string `json:"email`
	PDFPath string `json:"pdfPath"`
}

func SendInvoiceHandler(w http.ResponseWriter, r *http.Request) { //Define the function

	var req InvoiceRequest //Stores the request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil { //Handle JSON Errors
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer
	err = services.SendInvoiceEmail(req.Email, req.PDFPath)
	if err != nil { //Handle Email Errors
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Send Success Response
	response := map[string]string{
		"message": "Invoice sent successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

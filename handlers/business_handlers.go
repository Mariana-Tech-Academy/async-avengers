package handlers

import (
	"encoding/json"
	"invoiceSys/models"
	"invoiceSys/services"
	"net/http"

)

type BusinessHandler struct {
	Service *services.BusinessService
}

// Receives the request & creates a new bp
func (h *BusinessHandler) CreateBusiness(w http.ResponseWriter, r *http.Request) {
	
	// get user ID from token instead of request body
	userID := r.Context().Value("user_id").(uint)
	
	// collects business details from the request body
	var business models.Business
	err := json.NewDecoder(r.Body).Decode(&business)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set user ID from token automatically
	business.UserID = userID
	err = h.Service.CreateBusiness(&business)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	// return the created bp
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(business)
}

// Business details appear on invoices - retrieves the bp by user ID
func (h *BusinessHandler) GetBusiness(w http.ResponseWriter, r *http.Request) {
	// get user ID from token instead of URL
	userID := r.Context().Value("user_id").(uint)

	// US 1.1: fetch business profile so details can be used on invoices
	business, err := h.Service.GetBusiness(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(business)
}

// Edit business info - received requestes and updated business details
func (h *BusinessHandler) UpdateBusiness(w http.ResponseWriter, r *http.Request) {
	// get user ID from token instead of URL
	userID := r.Context().Value("user_id").(uint)

	// US 1.2 + US 1.3: collect updated business details from request body
	// (name, address, phone, email, logo, vat number, tax rate)
	var updated models.Business
	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer to update the bp
	err = h.Service.UpdateBusiness(uint(userID), &updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the updated bp
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)

}


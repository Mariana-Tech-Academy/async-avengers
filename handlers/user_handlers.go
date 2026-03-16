package handlers

import (
	"encoding/json"
	"invoiceSys/models"
	"invoiceSys/services"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	// collect request details
	var signUp models.User
	err := json.NewDecoder(r.Body).Decode(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer
	err = h.Service.RegisterUser(&signUp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(signUp)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// get login data from request body
	var login models.User
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer
	token, err := h.Service.Login(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *UserHandler) CreateCakeInvoice(w http.ResponseWriter, r *http.Request) {
	var cake models.CakeItem

	// 1. Decode the cake details (Size, Filler, Flavor, etc.)
	if err := json.NewDecoder(r.Body).Decode(&cake); err != nil {
		http.Error(w, "Invalid cake data", http.StatusBadRequest)
		return
	}

	// 2. Call the service layer to handle the math and saving
	// (You will need to add this method to your service file next!)
	if err := h.Service.SaveCakeInvoice(&cake); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cake)
}

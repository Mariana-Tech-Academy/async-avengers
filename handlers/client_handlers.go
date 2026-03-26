package handlers

import (
	"encoding/json"
	"invoiceSys/models"
	"invoiceSys/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ClientHandler struct {
	Service *services.ClientService
}

// US 2.1 - Add Client: receives the request and creates a new client
func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	// get user ID from token instead of request body
	userID := r.Context().Value("user_id").(uint)
	
	var client models.Client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set user ID from token automatically
	client.UserID = userID

	// call service layer to save the client
	err = h.Service.CreateClient(&client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the created client in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

// US 2.1 - Get client by ID: retrieves a single client
func (h *ClientHandler) GetClientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, err := strconv.Atoi(vars["clientID"])
	if err != nil {
		http.Error(w, "invalid client ID", http.StatusBadRequest)
		return
	}

	// get user ID from token
	userID := r.Context().Value("user_id").(uint)
	client, err := h.Service.GetClientByID(uint(clientID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// make sure the client belongs to the logged in user
	if client.UserID != userID {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

// US 2.3 - View Client Invoice History: retrieves all clients for a business owner
func (h *ClientHandler) GetClientsByUserID(w http.ResponseWriter, r *http.Request) {
	// get user ID from token instead of URL
	userID := r.Context().Value("user_id").(uint)
	clients, err := h.Service.GetClientsByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

// US 2.2 - Edit Client Details: receives the request and updates client details
func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, err := strconv.Atoi(vars["clientID"])
	if err != nil {
		http.Error(w, "invalid client ID", http.StatusBadRequest)
		return
	}

	// collect updated client details from rb
	var updated models.Client
	err = json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// make sure client belongs to logged in user
	userID := r.Context().Value("user_id").(uint)
	client, err := h.Service.GetClientByID(uint(clientID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if client.UserID != userID {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	err = h.Service.UpdateClient(uint(clientID), &updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the updated client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}
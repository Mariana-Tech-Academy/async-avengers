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

// 2.1: Create a new client
func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
    // collect client details from the rb
	var client models.Client
    if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	
	// call service layer to save the client
    if err := h.Service.CreateClient(&client); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// return the created client in the response
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

	client, err := h.Service.GetClientByID(uint(clientID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

// US 2.3 - View Client Invoice History: retrieves all clients for a business owner
func (h *ClientHandler) GetClientsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	clients, err := h.Service.GetClientsByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

//2.2 - edit client details
func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) { // 2.2 edit client details
    vars := mux.Vars(r)
    clientID, err := strconv.Atoi(vars["clientID"]) // Getting ID from URL
    if err != nil {
        http.Error(w, "invalid client ID", http.StatusBadRequest)
        return
    }

    var updated models.Client // take in new client detail from request body
	err = json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    // Call service layer to update the client in DB
    err = h.Service.UpdateClient(uint(clientID), &updated)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// return the updated client
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updated)
}


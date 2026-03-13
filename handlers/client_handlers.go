
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
    var client models.Client
    if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.Service.CreateClient(&client); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(client)
}
//2.2 - edit client details
func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) { // 2.2 edit client details
    vars := mux.Vars(r)
    clientID, err := strconv.Atoi(vars["id"]) // Getting ID from URL
    if err != nil {
        http.Error(w, "invalid client ID", http.StatusBadRequest)
        return
    }

    var updated models.Client // take in new client detail from request body
    if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call service layer to update the client in DB
    err = h.Service.UpdateClient(uint(clientID), &updated)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updated)
}

func (h *ClientHandler) GetClientInvoices(w http.ResponseWriter, r *http.Request) { // 2.3 viww client invoice history 
    vars := mux.Vars(r)
    clientID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "invalid client ID", http.StatusBadRequest)
        return
    }

    invoices, err := h.Service.GetInvoicesByClientID(uint(clientID)) // get invoice linked to client ID
    if err != nil {
        http.Error(w, "Could not find invoices for this client", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(invoices)
}
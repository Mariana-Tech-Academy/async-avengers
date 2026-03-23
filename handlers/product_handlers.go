package handlers

import (
    "encoding/json"
    "invoiceSys/models"
    "invoiceSys/services"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)

type ProductHandler struct {
    Service *services.ProductService
}

// US 3.1 – Create Product or Service
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    
    // Decode the request body
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Call service layer
    if err := h.Service.CreateProduct(&product); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

// US 3.2 – Edit Product or Service
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
    // Get product ID from URL params (e.g., /products/{id})
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    var updatedData models.Product
    if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Call service layer to handle the "Fetch -> Update -> Save" logic
    if err := h.Service.UpdateProduct(uint(id), &updatedData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
}
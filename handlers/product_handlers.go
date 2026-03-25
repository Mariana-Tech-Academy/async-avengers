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

// US 3.1 - Create Product - receives the request and creates a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// collect product details from the request body
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer to save the product
	err = h.Service.CreateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the created product in the response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// US 3.1 - Get product by ID - retrieves a single product
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["productID"])
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.Service.GetProductByID(uint(productID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// US 3.1 - Get all products - retrieves all products for a business owner
func (h *ProductHandler) GetProductsByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "invalid user ID", http.StatusBadRequest)
		return
	}

	products, err := h.Service.GetProductsByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// US 3.2 - Edit Product - receives the request and updates product details
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID, err := strconv.Atoi(vars["productID"])
	if err != nil {
		http.Error(w, "invalid product ID", http.StatusBadRequest)
		return
	}

	// collect updated product details from rb
	var updated models.Product
	err = json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer to update the product
	err = h.Service.UpdateProduct(uint(productID), &updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the updated product
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}
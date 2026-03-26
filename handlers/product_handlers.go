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

// Create Product - receives the request and creates a new product
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// get user ID from token instead of request body
	userID := r.Context().Value("user_id").(uint)

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// set user ID from token automatically
	product.UserID = userID

	// call service layer to save the product
	err = h.Service.CreateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the created product in the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// Get product by ID - retrieves a single product
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
	// make sure product belongs to logged in user
	userID := r.Context().Value("user_id").(uint)
	if product.UserID != userID {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// Get all products - retrieves all products for a business owner
func (h *ProductHandler) GetProductsByUserID(w http.ResponseWriter, r *http.Request) {
	// get user ID from token instead of URL
	userID := r.Context().Value("user_id").(uint)

	products, err := h.Service.GetProductsByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// Edit Product - receives the request and updates product details
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

	// make sure product belongs to logged in user
	userID := r.Context().Value("user_id").(uint)
	product, err := h.Service.GetProductByID(uint(productID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if product.UserID != userID {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	err = h.Service.UpdateProduct(uint(productID), &updated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return the updated product

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}

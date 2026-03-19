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

	//response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Epic 5.2 
func (h *UserHandler) DownloadInvoicePDF(c *gin.Context) {
	idParam := c.Param("id")

	invoiceID err := strconv.ParseUnit(idParam,10,64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid invoice"})
		return
	}
	pdfBytes, err := h.service.DownloadInvoicePDF(uint(invoiceID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	filename := "invoice_" + idParam + ".pdf"
	c.Header()

	func (r *UserRepo) GetInvoiceByID(invoiceID uint) (*models.Invoice, error) {
	var invoice models.Invoice

	err := db.DB.
		Preload("User").
		Preload("Client").
		Preload("Items").
		First(&invoice, invoiceID).Error
	if err != nil {
		return nil, err
	}

	return &invoice, nil
}
}

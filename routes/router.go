package routes

import (
	"fmt"
	"invoiceSys/handlers"
	"invoiceSys/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler,
	businessHandler *handlers.BusinessHandler,
	clientHandler *handlers.ClientHandler,
	productHandler *handlers.ProductHandler,
	invoiceHandler *handlers.InvoiceHandler,
	pdfHandler *handlers.PDFHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Backend is working!")
	}).Methods("GET")

	//public routes
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")

	// //sub router for protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// //authenticated routes

	protected.HandleFunc("/business", businessHandler.CreateBusiness).Methods("POST") // US 1.1 - Create bp
	protected.HandleFunc("/business", businessHandler.GetBusiness).Methods("GET")     // US 1.1 - Business details appear on invoices
	protected.HandleFunc("/business", businessHandler.UpdateBusiness).Methods("PUT")  // US 1.2 + US 1.3 - Edit business info & add tax info

	protected.HandleFunc("/clients", clientHandler.CreateClient).Methods("POST")            // US 2.1 -  Add client
	protected.HandleFunc("/clients/user", clientHandler.GetClientsByUserID).Methods("GET")  // US 2.3 - Get all clients for a user
	protected.HandleFunc("/clients/{clientID}", clientHandler.GetClientByID).Methods("GET") // US 2.1 - Get client by ID
	protected.HandleFunc("/clients/{clientID}", clientHandler.UpdateClient).Methods("PUT")  // US 2.2 - Update client

	protected.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")             // US 3.1 - Create product
	protected.HandleFunc("/products/user", productHandler.GetProductsByUserID).Methods("GET")   // US 3.1 - Get all products for a user
	protected.HandleFunc("/products/{productID}", productHandler.GetProductByID).Methods("GET") // US 3.1 - Get product by ID
	protected.HandleFunc("/products/{productID}", productHandler.UpdateProduct).Methods("PUT")  // US 3.2 - Update product

	protected.HandleFunc("/invoices", invoiceHandler.CreateInvoice).Methods("POST")                          // US 4.1 - Create invoice
	protected.HandleFunc("/invoices/user", invoiceHandler.GetInvoicesByUserID).Methods("GET")                // US 4.1 - Get all invoices for a user
	protected.HandleFunc("/invoices/{invoiceID}", invoiceHandler.GetInvoiceByID).Methods("GET")              // US 4.1 - Get invoice by ID
	protected.HandleFunc("/invoices/client/{clientID}", invoiceHandler.GetInvoicesByClientID).Methods("GET") // US 2.3 - Get all invoices for a client
	protected.HandleFunc("/invoices/{invoiceID}", invoiceHandler.UpdateInvoice).Methods("PUT")               // US 4.3 - Update draft invoice

	protected.HandleFunc("/invoices/{invoiceID}/pdf", pdfHandler.DownloadInvoicePDF).Methods("GET") // US 5.2 - Download invoice as PDF

	protected.HandleFunc("/invoices/{invoiceID}/status", invoiceHandler.UpdateInvoiceStatus).Methods("PUT") // US 6.1 - Update invoice status
	protected.HandleFunc("/invoices/{invoiceID}/status", invoiceHandler.GetInvoiceStatus).Methods("GET")    // US 6.2 - View invoice status

	return r

}

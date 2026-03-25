package routes

import (
	"invoiceSys/handlers"
	"invoiceSys/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, 
	businessHandler *handlers.BusinessHandler,
	clientHandler *handlers.ClientHandler,
	productHandler *handlers.ProductHandler,) *mux.Router {
	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")

	// //sub router for protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// //authenticated routes

	// US 1
	protected.HandleFunc("/business", businessHandler.CreateBusiness).Methods("POST") // Create bp - POST /business
	protected.HandleFunc("/business/{userID}", businessHandler.GetBusiness).Methods("GET") // Business details appear on invoices - GET /business/{userID}
	protected.HandleFunc("/business/{userID}", businessHandler.UpdateBusiness).Methods("PUT") // Edit business info & add tax info - PUT /business/{userID}

	// US 2
	protected.HandleFunc("/clients", clientHandler.CreateClient).Methods("POST") // Add client
	protected.HandleFunc("/clients/{clientID}", clientHandler.GetClientByID).Methods("GET") // Get client
	protected.HandleFunc("/clients/user/{userID}", clientHandler.GetClientsByUserID).Methods("GET") // Get all clients for a user
	protected.HandleFunc("/clients/{clientID}", clientHandler.UpdateClient).Methods("PUT") // Update client

	// US 3
	protected.HandleFunc("/products", productHandler.CreateProduct).Methods("POST") // Create product
	protected.HandleFunc("/products/{productID}", productHandler.GetProductByID).Methods("GET") // Get product by ID
	protected.HandleFunc("/products/user/{userID}", productHandler.GetProductsByUserID).Methods("GET") // Get all products for a use
	protected.HandleFunc("/products/{productID}", productHandler.UpdateProduct).Methods("PUT") // Edit Product


	return r

}

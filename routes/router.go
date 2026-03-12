package routes

import (
	"invoiceSys/handlers"
	"invoiceSys/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler, businessHandler *handlers.BusinessHandler) *mux.Router {
	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")

	// //sub router for protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// //authenticated routes

	// Create bp - POST /business
	protected.HandleFunc("/business", businessHandler.CreateBusiness).Methods("POST")
	// Business details appear on invoices - GET /business/{userID}
	protected.HandleFunc("/business/{userID}", businessHandler.GetBusiness).Methods("GET")
	// Edit business info & add tax info - PUT /business/{userID}
	protected.HandleFunc("/business/{userID}", businessHandler.UpdateBusiness).Methods("PUT")


	return r

}

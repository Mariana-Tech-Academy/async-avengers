package routes

import (
	"invoiceSys/handlers"
	"invoiceSys/middleware"

	"github.com/gorilla/mux"
)

func SetupRouter(userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	//public routes
	r.HandleFunc("/login", userHandler.Login).Methods("POST")
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")

	// //sub router for protected routes
	protected := r.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// //authenticated routes
	protected.HandleFunc("/user", userHandler.GetProfile).Methods("GET")     // US 1.1
	protected.HandleFunc("/user", userHandler.CreateProfile).Methods("POST") // US 1.1
	protected.HandleFunc("/user", userHandler.UpdateProfile).Methods("PUT")  // US 1.2

	return r

}

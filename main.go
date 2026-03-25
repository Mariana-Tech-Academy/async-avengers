package main

import (
	"fmt"
	"invoiceSys/db"
	"invoiceSys/handlers"
	"invoiceSys/repository"
	"invoiceSys/routes"
	"invoiceSys/services"
	"log"
	"net/http"
)

func main() {
	db.InitDb()

	// initialize repositories
	userRepo := &repository.UserRepo{}
	businessRepo := &repository.BusinessRepo{}
	clientRepo := &repository.ClientRepo{}
	productRepo := &repository.ProductRepo{}

	// initialize service
	userService := &services.UserService{Repo: userRepo}
	businessService := &services.BusinessService{Repo: businessRepo}
	clientService := &services.ClientService{Repo: clientRepo}
	productService := &services.ProductService{Repo: productRepo}

	// initialize handlers
	userHandler := &handlers.UserHandler{Service: userService}
	businessHandler := &handlers.BusinessHandler{Service: businessService}
	clientHandler := &handlers.ClientHandler{Service: clientService}
	productHandler := &handlers.ProductHandler{Service: productService}

	//routes
	r := routes.SetupRouter(userHandler, businessHandler, clientHandler, productHandler)

	//start server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
	fmt.Println("server started on :8080")
}

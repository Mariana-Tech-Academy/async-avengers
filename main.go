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
	invoiceRepo := &repository.InvoiceRepo{} // added invoice repo
	productRepo := &repository.ProductRepo{} // added product repo

	// initialize service
	userService := &services.UserService{Repo: userRepo}
	businessService := &services.BusinessService{Repo: businessRepo}
	clientService := &services.ClientService{Repo: clientRepo}
	invoiceService := &services.InvoiceService{Repo: invoiceRepo} // added invoice service
	productService := &services.ProductService{Repo: productRepo} // added product service

	// initialize handlers
	userHandler := &handlers.UserHandler{userService}
	businessHandler := &handlers.BusinessHandler{businessService}
	clientHandler := &handlers.ClientHandler{clientService}
	invoiceHandler := &handlers.InvoiceHandler{invoiceService} // added invoice handler
	productHandler := &handlers.ProductHandler{productService} // added product handler

	//routes
	r := routes.SetupRouter(userHandler, businessHandler, clientHandler, invoiceHandler, productHandler)  // added invoice and product handlers to routes

	//start server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
	fmt.Println("server started on :8080")
}

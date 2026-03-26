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
	"os"
)

func main() {
	db.InitDb()

	// initialize repositories
	userRepo := &repository.UserRepo{}
	businessRepo := &repository.BusinessRepo{}
	clientRepo := &repository.ClientRepo{}
	productRepo := &repository.ProductRepo{}
	invoiceRepo := &repository.InvoiceRepo{}

	// initialize services
	userService := &services.UserService{Repo: userRepo}
	businessService := &services.BusinessService{Repo: businessRepo}
	clientService := &services.ClientService{Repo: clientRepo}
	productService := &services.ProductService{Repo: productRepo}
	invoiceService := &services.InvoiceService{Repo:         invoiceRepo, BusinessRepo: businessRepo, ProductRepo:  productRepo,}

	// initialize handlers
	userHandler := &handlers.UserHandler{Service: userService}
	businessHandler := &handlers.BusinessHandler{Service: businessService}
	clientHandler := &handlers.ClientHandler{Service: clientService}
	productHandler := &handlers.ProductHandler{Service: productService}
	invoiceHandler := &handlers.InvoiceHandler{Service: invoiceService, ClientService: clientService,}
	pdfHandler := &handlers.PDFHandler{ InvoiceService:  invoiceService, BusinessService: businessService, ClientService:   clientService,}

	//routes
	r := routes.SetupRouter(userHandler, businessHandler, clientHandler, productHandler, invoiceHandler, pdfHandler)

	// start server
	port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
http.ListenAndServe(":"+port, r)
fmt.Println("server started on :8080")
	}

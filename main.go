package main

import (
	"fmt"
	"invoiceSys/db"
	"invoiceSys/handlers"
	"invoiceSys/repository"
	"invoiceSys/routes"
	"invoiceSys/services"
	"net/http"
	"os"
	"log"
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
// start server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("server started on :%s\n", port)
    err := http.ListenAndServe(":"+port, r)
    if err != nil {
        log.Fatal("failed to start server", err)
    }
} // This is the final closing brace for func main()
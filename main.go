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
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	
	db.InitDb()

	// initialize repositories
	userRepo := &repository.UserRepo{}
	businessRepo := &repository.BusinessRepo{}
	invoiceRepo := &repository.InvoiceRepo{}

	// initialize service
	userService := &services.UserService{Repo: userRepo}
	businessService := &services.BusinessService{Repo: businessRepo}
	invoiceService := &services.invoiceService{Repo:invoiceRepo}

	// initialize handlers
	userHandler := &handlers.UserHandler{userService}
	businessHandler := &handlers.BusinessHandler{businessService}
	invoiceHandler := &handlers.InvoiceHandler{invoiceService}


	//routes
	r := routes.SetupRouter(userHandler, businessHandler, invoiceHandler)

	//start server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
	fmt.Println("server started on :8080")
}

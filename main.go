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

	// initialize service
	userService := &services.UserService{Repo: userRepo}
	businessService := &services.BusinessService{Repo: businessRepo}
	clientService := &services.ClientService{Repo: clientRepo}

	// initialize handlers
	userHandler := &handlers.UserHandler{userService}
	businessHandler := &handlers.BusinessHandler{businessService}
	clientHandler := &handlers.ClientHandler{clientService}

	//routes
	r := routes.SetupRouter(userHandler, businessHandler, clientHandler)

	//start server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
	fmt.Println("server started on :8080")
}

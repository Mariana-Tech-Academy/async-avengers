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
	

	// initialize service
	userService := &services.UserService{Repo: userRepo}

	// initialize handlers
	userHandler := &handlers.UserHandler{userService}

	//routes
	r := routes.SetupRouter(userHandler)

	//start server
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("failed to start server", err)
	}
	fmt.Println("server started on :8080")
}

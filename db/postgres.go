package db

import (
	"fmt"
	"invoiceSys/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//connect to database
	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var dbErr error
	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if dbErr != nil {
		panic(dbErr)
	}

	//migrate the schema
	dbErr = DB.AutoMigrate(models.User{}, models.Business{})
	if dbErr != nil {
		log.Fatal("failed to migrate schema", dbErr)
	}

	fmt.Println("connected to database successfully!")
}

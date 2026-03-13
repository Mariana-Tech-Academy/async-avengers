package repository

import (
	"invoiceSys/db"
	"invoiceSys/models"
)

// Client Repository defines the database operations for clients
type ClientRepository interface {
	CreateClient(client *models.Client) error                    // US 2.1: saves a new client to the database
	GetClientByID(clientID uint) (*models.Client, error)         // US 2.1: retrieves a client by ID
	GetClientsByUserID(userID uint) ([]models.Client, error)     // US 2.3: retrieves all clients for a business owner
	UpdateClient(client *models.Client) error                    // US 2.2: updates existing client in the database
}

type ClientRepo struct{}

// US 2.1 - Add Client: saves the new client to the database
func (r *ClientRepo) CreateClient(client *models.Client) error {
	err := db.DB.Create(&client).Error
	if err != nil {
		return err
	}
	return nil
}

// US 2.1 - Get client by ID: fetches a single client from the database
func (r *ClientRepo) GetClientByID(clientID uint) (*models.Client, error) {
	var client models.Client
	err := db.DB.Where("id = ?", clientID).First(&client).Error
	if err != nil {
		return &models.Client{}, err
	}
	return &client, nil
}

// US 2.3 - View Client Invoice History: fetches all clients for a business owner
func (r *ClientRepo) GetClientsByUserID(userID uint) ([]models.Client, error) {
	var clients []models.Client
	err := db.DB.Where("user_id = ?", userID).Find(&clients).Error
	if err != nil {
		return []models.Client{}, err
	}
	return clients, nil
}

// US 2.2 - Edit Client Details: saves updated client details to the database
func (r *ClientRepo) UpdateClient(client *models.Client) error {
	err := db.DB.Save(&client).Error
	if err != nil {
		return err
	}
	return nil
}
package services

import (
	"errors"
	"invoiceSys/models"
	"invoiceSys/repository"
)

type ClientService struct {
	Repo repository.ClientRepository
}

// US 2.1 - Add Client: handles the logic for creating a new client
func (s *ClientService) CreateClient(client *models.Client) error {
	// save the new client to the database
	err := s.Repo.CreateClient(client)
	if err != nil {
		return err
	}
	return nil
}

// US 2.1 - Get client by ID: retrieves a single client
func (s *ClientService) GetClientByID(clientID uint) (*models.Client, error) {
	client, err := s.Repo.GetClientByID(clientID)
	if err != nil {
		return nil, errors.New("client not found")
	}
	return client, nil
}

// US 2.3 - View Client Invoice History: retrieves all clients for a business owner
func (s *ClientService) GetClientsByUserID(userID uint) ([]models.Client, error) {
	clients, err := s.Repo.GetClientsByUserID(userID)
	if err != nil {
		return nil, errors.New("no clients found")
	}
	return clients, nil
}

// US 2.2 - Edit Client Details: handles the logic for updating client details
func (s *ClientService) UpdateClient(clientID uint, updated *models.Client) error {
	existing, err := s.Repo.GetClientByID(clientID)
	if err != nil {
		return errors.New("client not found")
	}

	// US 2.2: update client details - changes apply to future invoices
	existing.Name = updated.Name
	existing.Email = updated.Email
	existing.Address = updated.Address

	// save the updated client
	err = s.Repo.UpdateClient(existing)
	if err != nil {
		return err
	}
	return nil
}
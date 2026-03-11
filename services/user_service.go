package services

import (
	"errors"
	"invoiceSys/middleware"
	"invoiceSys/models"
	"invoiceSys/repository"
	"invoiceSys/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) SaveCakeInvoice(item *models.CakeItem) error {
    // 1. Calculate base cost (Quantity * Unit Price)
    subtotal := float64(item.Quantity) * item.UnitPrice

    // 2. Add the Service Charge (delivery, rush, etc.)
    amountBeforeTax := subtotal + item.ServiceCharge

    // 3. Apply the Tax Rate
    taxAmount := amountBeforeTax * (item.TaxRate / 100)

    // 4. Set the final Total
    item.Total = amountBeforeTax + taxAmount

    // 5. Save to database using your repository
    return s.Repo.CreateCakeItem(item) 
}

func (s *UserService) RegisterUser(req *models.User) error {

	// Check if user already exists
	_, err := s.Repo.GetUserByUsername(req.Username)
	if err == nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = hashedPass

	// Save user to DB
	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(req *models.User) (string, error) {

	// Check if user exists
	user, err := s.Repo.GetUserByUsername(req.Username)
	if err != nil {
		return "", err
	}

	// Compare password
	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Generate JWT token
	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

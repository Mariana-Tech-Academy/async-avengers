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

const TaxRate = 0.075

func (s *UserService) SaveCakeInvoice(item *models.CakeItem) error {

	subtotal := float64(item.Quantity) * item.UnitPrice // 1. Calculate base cost (Quantity * Unit Price)

	amountBeforeTax := subtotal + item.ServiceCharge // 2. Add the Service Charge

	// Use the constant you defined on line 14
	taxAmount := amountBeforeTax * TaxRate   // already chnged in user.go 3. Apply the Tax Rate
	item.TaxAmount = taxAmount               //maee sure tax amount is stated even if corretly calculated
	item.Total = amountBeforeTax + taxAmount // 4. Set the final Total

	return s.Repo.CreateCakeItem(item) // 5. Save to database using your repository
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

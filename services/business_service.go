package services
import (
	"errors"
	"invoiceSys/models"
	"invoiceSys/repository"
	
)

type BusinessService struct {
	Repo repository.BusinessRepository
}

// US 1.1 Create business profile - handles the logic for it
func (s *BusinessService) CreateBusiness(business *models.Business) error {

	// checks if a business profile already exists for this user
	_, err := s.Repo.GetBusinessByUserID(business.UserID)
	if err == nil {
		return errors.New("Error - Business profile already exists!")
	}

	// saves the new business profile
	err = s.Repo.CreateBusiness(business)
	if err != nil {
		return err
	}
	return nil
}

// Business details appear on invoices - gets bp by user ID
func (s *BusinessService) GetBusiness(userID uint) (*models.Business, error) {
	business, err := s.Repo.GetBusinessByUserID(userID)
	if err != nil {
		return nil, errors.New("business profile not found")
	}
	return business, nil
}

// Edit business info - handles the logic of updating business details
func (s *BusinessService) UpdateBusiness(userID uint, updated *models.Business) error {
	// fetching existing bp before updating
	existing, err := s.Repo.GetBusinessByUserID(userID)
	if err != nil {
		return errors.New("Error - Business profile not found!")
	}

	// Update business details
	existing.Name = updated.Name
	existing.Address = updated.Address
	existing.Phone = updated.Phone
	existing.Email = updated.Email
	existing.Logo = updated.Logo

	// Add Tax info - update VAT no. & tax rate
	existing.VATNumber = updated.VATNumber
	existing.TaxRate = updated.TaxRate

	// Save the updated details
	err = s.Repo.UpdateBusiness(existing)
	if err != nil {
		return err
	}
	return nil
}

// US 1.3 - Tax amount automatically calculated based on tax rate
func (s *BusinessService) CalculateTax(amount float64, userID uint) (float64, error) {
	// fetching the bp to get the tax rate
	business, err := s.Repo.GetBusinessByUserID(userID)
	if err != nil {
		return 0, errors.New("Error - Business profile not found!")
	}

	// Automatically calculates tax amount
	tax := amount * (business.TaxRate / 100)
	return tax, nil
}
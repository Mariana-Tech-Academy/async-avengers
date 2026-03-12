package repository

import (
	"invoiceSys/db"
	"invoiceSys/models"
)

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
	CreateCakeItem(item *models.CakeItem) error //added command

	CreateClient(client *models.Client) error              // Added for US 2.1 mx client
	UpdateClient(client *models.Client) error              // Added for US 2.2 mx client
	GetInvoicesByClientID(clientID uint) ([]models.Invoice, error) // Added for US 2.3 mx client


	CreateInvoice(invoice *models.Invoice) error // invoice/items
}


type UserRepo struct{}

// Handler layer-------->services layer-------->repository layer(db methods)









func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := db.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepo) CreateCakeItem(item *models.CakeItem) error {
    return db.DB.Create(item).Error // save cake order details 
}


package repository

import (
	//"invoiceSys/db" dependency injection using r*

	"invoiceSys/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
	UpdateUser(user *models.User) error

	CreateClient(client *models.Client) error // 2,1 mx client
	UpdateClient(client *models.Client) error //2.2 mx client

	GetInvoicesByClientID(clientID uint) ([]models.Invoice, error) // 2.3 mx client
	CreateInvoice(invoice *models.Invoice) error                   //invoice items

	CreateCakeItem(item *models.CakeItem) error //added command
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
}

type UserRepo struct {
	DB *gorm.DB
}

// Handler layer-------->services layer-------->repository layer(db methods)

func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

func (r *UserRepo) CreateUser(user *models.User) error {
	err := r.DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UpdateUser(user *models.User) error { // 1.2 update business profile
	return r.DB.Save(user).Error
}

func (r *UserRepo) CreateClient(client *models.Client) error {
	return r.DB.Create(client).Error // 2.1 create clieng
}

func (r *UserRepo) UpdateClient(client *models.Client) error { // 2.2 edit client
	return r.DB.Create(client).Error
}

func (r *UserRepo) CreateInvoice(invoice *models.Invoice) error { // 1.1 save cake items+ create invoice
	return r.DB.Create(invoice).Error
}

func (r *UserRepo) CreateProduct(product *models.Product) error { // US 3.1 - Create Product
	return r.DB.Create(product).Error
}

func (r *UserRepo) UpdateProduct(product *models.Product) error {
	return r.DB.Save(product).Error // 3.2 update product + changes to name/price/description
}

func (r *UserRepo) CreateCakeItem(item *models.CakeItem) error {
	return r.DB.Create(item).Error
}

func (r *UserRepo) GetInvoicesByClientID(clientID uint) ([]models.Invoice, error) { //2.3 client invoice history to view
	var invoices []models.Invoice
	err := r.DB.Preload("Items").Where("client_id = ?", clientID).Find(&invoices).Error // get cake items assoc with specific invoicw
	return invoices, err
}

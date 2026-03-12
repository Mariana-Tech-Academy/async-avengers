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


type UserRepo struct{

	DB *gorm.DB
}

// Handler layer-------->services layer-------->repository layer(db methods)



func (r *UserRepo) UpdateUser(user *models.User) error { // 1.2 update business profile
	return db.DB.Save(user).Error
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


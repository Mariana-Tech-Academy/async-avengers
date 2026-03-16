package repository

import (
	//"invoiceSys/db" dependency injection using r*

	"invoiceSys/models"
	"gorm.io/gorm"
)

type ProductRepository interface {

	CreateCakeItem(item *models.CakeItem) error //added command
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
}

type ProductRepo struct {
	DB *gorm.DB
}


func (r *ProductRepo) CreateCakeItem(item *models.CakeItem) error {
	return r.DB.Create(item).Error
}

func (r *ProductRepo) CreateProduct(product *models.Product) error { // US 3.1 - Create Product
	return r.DB.Create(product).Error
}

func (r *ProductRepo) UpdateProduct(product *models.Product) error {
	return r.DB.Save(product).Error // 3.2 update product + changes to name/price/description
}

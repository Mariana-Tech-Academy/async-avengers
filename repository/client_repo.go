
package repository

import (
	//"invoiceSys/db" dependency injection using r*

	"invoiceSys/models"
	"gorm.io/gorm"
)


type clientRepository interface {
	CreateClient(client *models.Client) error // 2,1 mx client
	UpdateClient(client *models.Client) error //2.2 mx client
}


type ClientRepo struct {
	DB *gorm.DB
}

func (r *ClientRepo) CreateClient(client *models.Client) error {
	return r.DB.Create(client).Error // 2.1 create clieng
}

func (r *ClientRepo) UpdateClient(client *models.Client) error { // 2.2 edit client
	return r.DB.Create(client).Error
}

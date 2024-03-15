package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type AddressRepository interface {
	Create(model *model.Address, tx *gorm.DB) (*model.Address, error)
	Update(model *model.Address, tx *gorm.DB) error
	Delete(model *model.Address, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Address, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Address, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Address) error
}

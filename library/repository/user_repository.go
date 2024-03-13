package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type UserRepository interface {
	Create(model *model.User, tx *gorm.DB) (*model.User, error)
	Update(model *model.User, tx *gorm.DB) error
	Delete(model *model.User, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.User, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.User, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.User) error
}

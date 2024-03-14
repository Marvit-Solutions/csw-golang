package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ClassUserRepository interface {
	Create(model *model.ClassUser, tx *gorm.DB) (*model.ClassUser, error)
	Update(model *model.ClassUser, tx *gorm.DB) error
	Delete(model *model.ClassUser, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ClassUser, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ClassUser, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ClassUser) error
}

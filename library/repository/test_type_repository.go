package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type TestTypeRepository interface {
	Create(model *model.TestType, tx *gorm.DB) (*model.TestType, error)
	Update(model *model.TestType, tx *gorm.DB) error
	Delete(model *model.TestType, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.TestType, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.TestType, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.TestType) error
}

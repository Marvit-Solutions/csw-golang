package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ModuleRepository interface {
	Create(model *model.Module, tx *gorm.DB) (*model.Module, error)
	Update(model *model.Module, tx *gorm.DB) error
	Delete(model *model.Module, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Module, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Module, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Module) error
}

package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type SubModuleRepository interface {
	Create(model *model.SubModule, tx *gorm.DB) (*model.SubModule, error)
	Update(model *model.SubModule, tx *gorm.DB) error
	Delete(model *model.SubModule, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.SubModule, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.SubModule, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.SubModule) error
}

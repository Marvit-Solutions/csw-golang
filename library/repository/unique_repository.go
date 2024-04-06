package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type UniqueRepository interface {
	Create(model *model.Unique, tx *gorm.DB) (*model.Unique, error)
	Update(model *model.Unique, tx *gorm.DB) error
	Delete(model *model.Unique, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Unique, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Unique, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Unique) error
}

package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type MediaRepository interface {
	Create(model *model.Media, tx *gorm.DB) (*model.Media, error)
	Update(model *model.Media, tx *gorm.DB) error
	Delete(model *model.Media, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Media, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Media, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Media) error
}

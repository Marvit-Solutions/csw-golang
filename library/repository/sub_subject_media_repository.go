package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type SubSubjectMediaRepository interface {
	Create(model *model.SubSubjectMedia, tx *gorm.DB) (*model.SubSubjectMedia, error)
	Update(model *model.SubSubjectMedia, tx *gorm.DB) error
	Delete(model *model.SubSubjectMedia, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.SubSubjectMedia, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.SubSubjectMedia, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.SubSubjectMedia) error
}

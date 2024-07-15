package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type SubSubjectMediumRepository interface {
	Create(model *model.SubSubjectMedium, tx *gorm.DB) (*model.SubSubjectMedium, error)
	Update(model *model.SubSubjectMedium, tx *gorm.DB) error
	Delete(model *model.SubSubjectMedium, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.SubSubjectMedium, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.SubSubjectMedium, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.SubSubjectMedium) error
}

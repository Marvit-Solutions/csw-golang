package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type SubSubjectRepository interface {
	Create(model *model.SubSubject, tx *gorm.DB) (*model.SubSubject, error)
	Update(model *model.SubSubject, tx *gorm.DB) error
	Delete(model *model.SubSubject, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.SubSubject, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.SubSubject, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.SubSubject) error
}

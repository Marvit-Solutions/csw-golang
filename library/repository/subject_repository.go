package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type SubjectRepository interface {
	Create(model *model.Subject, tx *gorm.DB) (*model.Subject, error)
	Update(model *model.Subject, tx *gorm.DB) error
	Delete(model *model.Subject, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Subject, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Subject, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Subject) error
}

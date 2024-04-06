package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ExerciseRepository interface {
	Create(model *model.Exercise, tx *gorm.DB) (*model.Exercise, error)
	Update(model *model.Exercise, tx *gorm.DB) error
	Delete(model *model.Exercise, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Exercise, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Exercise, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Exercise) error
}

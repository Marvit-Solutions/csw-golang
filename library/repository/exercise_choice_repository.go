package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ExerciseChoiceRepository interface {
	Create(model *model.ExerciseChoice, tx *gorm.DB) (*model.ExerciseChoice, error)
	Update(model *model.ExerciseChoice, tx *gorm.DB) error
	Delete(model *model.ExerciseChoice, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ExerciseChoice, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseChoice, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ExerciseChoice) error
}

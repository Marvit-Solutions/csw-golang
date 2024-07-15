package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ExerciseSubmissionsModuleRepository interface {
	Create(model *model.ExerciseSubmissionsModule, tx *gorm.DB) (*model.ExerciseSubmissionsModule, error)
	Update(model *model.ExerciseSubmissionsModule, tx *gorm.DB) error
	Delete(model *model.ExerciseSubmissionsModule, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ExerciseSubmissionsModule, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseSubmissionsModule, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ExerciseSubmissionsModule) error
}

package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ExerciseSubmissionsSubModuleRepository interface {
	Create(model *model.ExerciseSubmissionsSubModule, tx *gorm.DB) (*model.ExerciseSubmissionsSubModule, error)
	Update(model *model.ExerciseSubmissionsSubModule, tx *gorm.DB) error
	Delete(model *model.ExerciseSubmissionsSubModule, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ExerciseSubmissionsSubModule, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseSubmissionsSubModule, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ExerciseSubmissionsSubModule) error
}

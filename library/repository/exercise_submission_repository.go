package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ExerciseSubmissionRepository interface {
	Create(model *model.ExerciseSubmission, tx *gorm.DB) (*model.ExerciseSubmission, error)
	Update(model *model.ExerciseSubmission, tx *gorm.DB) error
	Delete(model *model.ExerciseSubmission, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ExerciseSubmission, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseSubmission, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ExerciseSubmission) error
}

package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ExerciseQuestionRepository interface {
	Create(model *model.ExerciseQuestion, tx *gorm.DB) (*model.ExerciseQuestion, error)
	Update(model *model.ExerciseQuestion, tx *gorm.DB) error
	Delete(model *model.ExerciseQuestion, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ExerciseQuestion, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseQuestion, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ExerciseQuestion) error
}

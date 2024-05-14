package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

// Defines the interface for model repository operations.
type ExerciseQuestionMediaRepository interface {
	Create(model *model.ExerciseQuestionMedia, tx *gorm.DB) (*model.ExerciseQuestionMedia, error)
	Update(model *model.ExerciseQuestionMedia, tx *gorm.DB) error
	Delete(model *model.ExerciseQuestionMedia, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ExerciseQuestionMedia, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseQuestionMedia, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ExerciseQuestionMedia) error
}

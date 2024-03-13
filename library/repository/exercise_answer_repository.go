package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ExerciseAnswerRepository interface {
	Create(model *model.ExerciseAnswer, tx *gorm.DB) (*model.ExerciseAnswer, error)
	Update(model *model.ExerciseAnswer, tx *gorm.DB) error
	Delete(model *model.ExerciseAnswer, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ExerciseAnswer, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseAnswer, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ExerciseAnswer) error
}

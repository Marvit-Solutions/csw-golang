package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ExerciseScoreRepository interface {
	Create(model *model.ExerciseScore, tx *gorm.DB) (*model.ExerciseScore, error)
	Update(model *model.ExerciseScore, tx *gorm.DB) error
	Delete(model *model.ExerciseScore, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ExerciseScore, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ExerciseScore, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ExerciseScore) error
}

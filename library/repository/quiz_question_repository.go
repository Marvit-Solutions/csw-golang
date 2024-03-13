package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type QuizQuestionRepository interface {
	Create(model *model.QuizQuestion, tx *gorm.DB) (*model.QuizQuestion, error)
	Update(model *model.QuizQuestion, tx *gorm.DB) error
	Delete(model *model.QuizQuestion, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.QuizQuestion, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.QuizQuestion, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.QuizQuestion) error
}

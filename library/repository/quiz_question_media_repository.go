package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type QuizQuestionMediaRepository interface {
	Create(model *model.QuizQuestionMedia, tx *gorm.DB) (*model.QuizQuestionMedia, error)
	Update(model *model.QuizQuestionMedia, tx *gorm.DB) error
	Delete(model *model.QuizQuestionMedia, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.QuizQuestionMedia, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.QuizQuestionMedia, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.QuizQuestionMedia) error
}

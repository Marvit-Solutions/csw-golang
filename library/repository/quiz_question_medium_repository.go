package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type QuizQuestionMediumRepository interface {
	Create(model *model.QuizQuestionMedium, tx *gorm.DB) (*model.QuizQuestionMedium, error)
	Update(model *model.QuizQuestionMedium, tx *gorm.DB) error
	Delete(model *model.QuizQuestionMedium, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.QuizQuestionMedium, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.QuizQuestionMedium, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.QuizQuestionMedium) error
}

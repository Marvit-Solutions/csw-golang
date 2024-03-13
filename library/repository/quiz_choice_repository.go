package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type QuizChoiceRepository interface {
	Create(model *model.QuizChoice, tx *gorm.DB) (*model.QuizChoice, error)
	Update(model *model.QuizChoice, tx *gorm.DB) error
	Delete(model *model.QuizChoice, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.QuizChoice, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.QuizChoice, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.QuizChoice) error
}

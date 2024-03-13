package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type QuizAnswerRepository interface {
	Create(model *model.QuizAnswer, tx *gorm.DB) (*model.QuizAnswer, error)
	Update(model *model.QuizAnswer, tx *gorm.DB) error
	Delete(model *model.QuizAnswer, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.QuizAnswer, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.QuizAnswer, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.QuizAnswer) error
}

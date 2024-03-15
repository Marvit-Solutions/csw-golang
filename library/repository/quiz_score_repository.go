package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type QuizScoreRepository interface {
	Create(model *model.QuizScore, tx *gorm.DB) (*model.QuizScore, error)
	Update(model *model.QuizScore, tx *gorm.DB) error
	Delete(model *model.QuizScore, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.QuizScore, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.QuizScore, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.QuizScore) error
}

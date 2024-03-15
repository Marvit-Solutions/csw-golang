package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type QuizSubmissionRepository interface {
	Create(model *model.QuizSubmission, tx *gorm.DB) (*model.QuizSubmission, error)
	Update(model *model.QuizSubmission, tx *gorm.DB) error
	Delete(model *model.QuizSubmission, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.QuizSubmission, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.QuizSubmission, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.QuizSubmission) error
}

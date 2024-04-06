package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type MentorRepository interface {
	Create(model *model.Mentor, tx *gorm.DB) (*model.Mentor, error)
	Update(model *model.Mentor, tx *gorm.DB) error
	Delete(model *model.Mentor, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Mentor, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Mentor, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Mentor) error
}

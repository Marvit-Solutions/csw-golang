package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type UserMentorTestimonialRepository interface {
	Create(model *model.UserMentorTestimonial, tx *gorm.DB) (*model.UserMentorTestimonial, error)
	Update(model *model.UserMentorTestimonial, tx *gorm.DB) error
	Delete(model *model.UserMentorTestimonial, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.UserMentorTestimonial, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.UserMentorTestimonial, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.UserMentorTestimonial) error
}

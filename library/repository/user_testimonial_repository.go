package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type UserTestimonialRepository interface {
	Create(model *model.UserTestimonial, tx *gorm.DB) (*model.UserTestimonial, error)
	Update(model *model.UserTestimonial, tx *gorm.DB) error
	Delete(model *model.UserTestimonial, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.UserTestimonial, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.UserTestimonial, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.UserTestimonial) error
}

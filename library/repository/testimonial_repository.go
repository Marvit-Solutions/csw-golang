package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type TestimonialRepository interface {
	Create(model *model.Testimonial, tx *gorm.DB) (*model.Testimonial, error)
	Update(model *model.Testimonial, tx *gorm.DB) error
	Delete(model *model.Testimonial, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Testimonial, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Testimonial, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Testimonial) error
}

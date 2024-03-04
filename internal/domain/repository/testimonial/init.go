package testimonial

import (
	"csw-golang/internal/domain/entity/dto"

	"gorm.io/gorm"
)

type TestimonialRepo interface {
	GetAllTestimonials() ([]dto.Testimonials, error)
}

type testimonialRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) TestimonialRepo {
	return &testimonialRepo{
		db,
	}
}

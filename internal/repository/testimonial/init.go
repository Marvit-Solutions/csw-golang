package testimonial

import (
	dto "csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type TestimonialRepository interface {
	GetAllTestimonials() ([]dto.Testimonials, error)
}

type testimonialRepository struct {
	db *gorm.DB
}

func NewTestimonialRepository(db *gorm.DB) TestimonialRepository {
	return &testimonialRepository{db}
}

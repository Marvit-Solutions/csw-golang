package testimonial

import (
	dto "csw-golang/internal/domain/response"
	"csw-golang/internal/repository/testimonial"

	"gorm.io/gorm"
)

type TestimonialUsecase interface {
	GetAllTestimonials() ([]dto.Testimonials, error)
}

type testimonialUsecase struct {
	testimonialRepo testimonial.TestimonialRepository
}

func NewTestimonialUsecase(
	db *gorm.DB,
) TestimonialUsecase {
	return &testimonialUsecase{
		testimonialRepo: testimonial.NewTestimonialRepository(db),
	}
}

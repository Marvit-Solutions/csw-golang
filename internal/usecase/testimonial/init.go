package testimonial

import (
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"
	"github.com/Marvit-Solutions/csw-golang/internal/repository/testimonial"

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

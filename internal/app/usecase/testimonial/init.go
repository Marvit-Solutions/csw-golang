package testimonial

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/repository/testimonial"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type TestimonialUsecase interface {
	GetAllTestimonials() ([]response.Testimonials, error)
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

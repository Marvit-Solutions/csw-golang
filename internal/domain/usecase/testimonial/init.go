package testimonial

import (
	"csw-golang/internal/domain/entity/dto"
	tr "csw-golang/internal/domain/repository/testimonial"
)

type TestimonialUsecase interface {
	GetListTopSixTestimonials() (error, dto.Testimonials)
	GetAllTestimonials() (error, dto.Testimonials)
}

type testimonialUsecase struct {
	testimonialRepo tr.TestimonialRepo
}

func New(testimonialRepo tr.TestimonialRepo) *testimonialUsecase {
	return &testimonialUsecase{
		testimonialRepo,
	}
}

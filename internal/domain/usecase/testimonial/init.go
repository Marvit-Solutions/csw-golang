package testimonial

import (
	"csw-golang/internal/domain/entity/dto"
	tr "csw-golang/internal/domain/repository/testimonial"
)

type TestimonialUsecase interface {
	GetListTopSixTestimonials() (error, dto.ListTestimonials)
	GetAllTestimonials() (error, dto.ListTestimonials)
}

type testimonialUsecase struct {
	testimonialRepo tr.TestimonialRepo
}

func New(testimonialRepo tr.TestimonialRepo) *testimonialUsecase {
	return &testimonialUsecase{
		testimonialRepo,
	}
}

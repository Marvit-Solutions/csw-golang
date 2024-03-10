package testimonial

import (
	tc "csw-golang/internal/domain/usecase/testimonial"
)

type TestimonialHandler struct {
	testimonialUsecase tc.TestimonialUsecase
}

func New(testimonialUsecase tc.TestimonialUsecase) *TestimonialHandler {
	return &TestimonialHandler{
		testimonialUsecase,
	}
}

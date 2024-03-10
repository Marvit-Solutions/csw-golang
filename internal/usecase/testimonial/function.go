package testimonial

import (
	dto "csw-golang/internal/domain/response"
)

func (tr *testimonialUsecase) GetAllTestimonials() ([]dto.Testimonials, error) {
	data, err := tr.testimonialRepo.GetAllTestimonials()
	if err != nil {
		return nil, err
	}
	return data, err
}

package testimonial

import (
	"csw-golang/internal/domain/entity/dto"
)

func (tr *testimonialUsecase) GetAllTestimonials() ([]dto.Testimonials, error) {
	data, err := tr.testimonialRepo.GetAllTestimonials()
	if err != nil {
		return nil, err
	}
	return data, err
}

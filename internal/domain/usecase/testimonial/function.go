package testimonial

import (
	"csw-golang/internal/domain/entity/dto"
)

func (tr *testimonialUsecase) GetListTopSixTestimonials() (error, dto.Testimonials) {
	err, mentor := tr.testimonialRepo.GetListTopSixTestimonials()
	if err != nil {
		return err, dto.Testimonials{}
	}
	return nil, mentor
}

func (tr *testimonialUsecase) GetAllTestimonials() (error, dto.Testimonials) {
	err, mentor := tr.testimonialRepo.GetAllTestimonials()
	if err != nil {
		return err, dto.Testimonials{}
	}
	return nil, mentor
}

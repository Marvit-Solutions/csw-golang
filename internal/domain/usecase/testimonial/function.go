package testimonial

import (
	"csw-golang/internal/domain/entity/dto"
)

func (tr *testimonialUsecase) GetListTopSixTestimonials() (error, dto.ListTestimonials) {
	err, mentor := tr.testimonialRepo.GetListTopSixTestimonials()
	if err != nil {
		return err, dto.ListTestimonials{}
	}
	return nil, mentor
}

func (tr *testimonialUsecase) GetAllTestimonials() (error, dto.ListTestimonials) {
	err, mentor := tr.testimonialRepo.GetAllTestimonials()
	if err != nil {
		return err, dto.ListTestimonials{}
	}
	return nil, mentor
}

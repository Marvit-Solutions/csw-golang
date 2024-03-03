package testimonial

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
)

func (t testimonialRepo) GetAllTestimonials() (error, dto.Testimonials) {

	var allTestimonials dto.Testimonials
	var testimonials []datastruct.Testimonials

	tx := t.db.Preload("UserDetail").Find(&testimonials)
	if tx.Error != nil {
		return tx.Error, allTestimonials
	}

	for _, testimonial := range testimonials {
		dtoTestimonial := struct {
			ID      string  `json:"ID"`
			Comment string  `json:"Comment"`
			Rating  float32 `json:"Rating"`
			User    struct {
				Name           string `json:"Name" form:"Name"`
				ProfilePicture string `json:"ProfilePicture" form:"ProfilePicture"`
				Class          string `json:"Class " form:"Class "`
			}
		}{
			ID:      testimonial.ID,
			Rating:  testimonial.Rating,
			Comment: testimonial.Comment,
			User: struct {
				Name           string `json:"Name" form:"Name"`
				ProfilePicture string `json:"ProfilePicture" form:"ProfilePicture"`
				Class          string `json:"Class " form:"Class "`
			}{
				Name:           testimonial.UserDetail.Name,
				ProfilePicture: testimonial.UserDetail.ProfilePicture,
				Class:          testimonial.UserDetail.Class,
			},
		}
		allTestimonials = append(allTestimonials, dtoTestimonial)
	}
	return nil, allTestimonials

}

package testimonial

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
)

func (t testimonialRepo) GetListTopSixTestimonials() (error, dto.Testimonials) {
	var topTestimonials dto.Testimonials
	var testimonials []datastruct.Testimonials

	// Fetch top six testimonials from the database
	tx := t.db.Order("rating desc").Limit(6).Find(&testimonials)
	if tx.Error != nil {
		return tx.Error, topTestimonials
	}
	for _, testimonial := range testimonials {
		dtoTestimonial := struct {
			ID      string  `json:"ID"`
			Comment string  `json:"comment"`
			Rating  float32 `json:"rating"`
			User    struct {
				Name           string `json:"Name" form:"Name"`
				ProfilePicture string `json:"ProfilePicture" form:"ProfilePicture"`
				Status         string `json:"Status" form:"Status"`
			}
		}{
			ID:     testimonial.ID,
			Rating: testimonial.Rating,
			User: struct {
				Name           string `json:"Name" form:"Name"`
				ProfilePicture string `json:"ProfilePicture" form:"ProfilePicture"`
				Status         string `json:"Status" form:"Status"`
			}{
				Name:           testimonial.UserDetail.Name,
				ProfilePicture: testimonial.UserDetail.ProfilePicture,
				Status:         testimonial.UserDetail.Status,
			},
		}
		topTestimonials = append(topTestimonials, dtoTestimonial)
	}

	return nil, topTestimonials
}

func (t testimonialRepo) GetAllTestimonials() (error, dto.Testimonials) {

	var allTestimonials dto.Testimonials
	var testimonials []datastruct.Testimonials

	tx := t.db.Find(&testimonials)
	if tx.Error != nil {
		return tx.Error, allTestimonials
	}

	for _, testimonial := range testimonials {
		dtoTestimonial := struct {
			ID      string  `json:"ID"`
			Comment string  `json:"comment"`
			Rating  float32 `json:"rating"`
			User    struct {
				Name           string `json:"Name" form:"Name"`
				ProfilePicture string `json:"ProfilePicture" form:"ProfilePicture"`
				Status         string `json:"Status" form:"Status"`
			}
		}{
			ID:     testimonial.ID,
			Rating: testimonial.Rating,
			User: struct {
				Name           string `json:"Name" form:"Name"`
				ProfilePicture string `json:"ProfilePicture" form:"ProfilePicture"`
				Status         string `json:"Status" form:"Status"`
			}{
				Name:           testimonial.UserDetail.Name,
				ProfilePicture: testimonial.UserDetail.ProfilePicture,
				Status:         testimonial.UserDetail.Status,
			},
		}
		allTestimonials = append(allTestimonials, dtoTestimonial)
	}
	return nil, allTestimonials

}

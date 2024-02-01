package testimonial

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
)

func (t testimonialRepo) GetListTopSixTestimonials() (error, dto.ListTestimonials) {
	var topTestimonials dto.ListTestimonials
	var testimonials []datastruct.Testimonials

	// Fetch top six testimonials from the database
	tx := t.db.Order("rating desc").Limit(6).Find(&testimonials)
	if tx.Error != nil {
		return tx.Error, topTestimonials
	}
	for _, testimonial := range testimonials {
		dtoTestimonial := struct {
			ID             string  `json:"ID"`
			Name           string  `json:"name"`
			Status         string  `json:"status"`
			ProfilePicture string  `json:"profile_photo"`
			Comment        string  `json:"comment"`
			Rating         float32 `json:"rating"`
		}{
			ID:             testimonial.ID,
			Name:           testimonial.Name,
			Status:         testimonial.Status,
			ProfilePicture: testimonial.ProfilePicture,
			Comment:        testimonial.Comment,
			Rating:         testimonial.Rating,
			// Add other fields as needed
		}
		topTestimonials = append(topTestimonials, dtoTestimonial)
	}

	return nil, topTestimonials
}

func (t testimonialRepo) GetAllTestimonials() (error, dto.ListTestimonials) {

	var allTestimonials dto.ListTestimonials
	var testimonials []datastruct.Testimonials

	// Fetch all testimonials from the database
	tx := t.db.Find(&testimonials)
	if tx.Error != nil {
		return tx.Error, allTestimonials
	}

	for _, testimonial := range testimonials {
		dtoTestimonial := struct {
			ID             string  `json:"ID"`
			Name           string  `json:"name"`
			Status         string  `json:"status"`
			ProfilePicture string  `json:"profile_photo"`
			Comment        string  `json:"comment"`
			Rating         float32 `json:"rating"`
		}{
			ID:             testimonial.ID,
			Name:           testimonial.Name,
			Status:         testimonial.Status,
			ProfilePicture: testimonial.ProfilePicture,
			Comment:        testimonial.Comment,
			Rating:         testimonial.Rating,
			// Add other fields as needed
		}
		allTestimonials = append(allTestimonials, dtoTestimonial)
	}
	return nil, allTestimonials

}

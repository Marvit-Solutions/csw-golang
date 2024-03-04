package testimonial

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
)

func (t testimonialRepo) GetAllTestimonials() ([]dto.Testimonials, error) {

	var testimonials []datastruct.Testimonials
	var users []datastruct.Users

	tx := t.db.Preload("UserTestimonials").Find(&testimonials)
	if tx.Error != nil {
		return nil, tx.Error
	}

	allTestimonials := make([]dto.Testimonials, 0)
	for _, testimonial := range testimonials {
		tx := t.db.Preload("UserDetails").
			Where("users.id = ?", testimonial.UserTestimonials.UserID).
			First(&users)
		if tx.Error != nil {
			return nil, tx.Error
		}

		newTestimonials := dto.Testimonials{
			ID:      testimonial.ID,
			Comment: testimonial.Comment,
			Rating:  testimonial.Rating,
			User: dto.UserTestimonialResponse{
				Name:           users[0].UserDetails.Name,
				ProfilePicture: users[0].UserDetails.ProfilePicture,
				Class:          users[0].UserDetails.Class,
			},
		}
		allTestimonials = append(allTestimonials, newTestimonials)
	}

	return allTestimonials, nil
}

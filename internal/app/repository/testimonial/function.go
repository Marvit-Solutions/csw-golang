package testimonial

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/datastruct"
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"
)

func (t testimonialRepository) GetAllTestimonials() ([]dto.Testimonials, error) {
	var testimonials []datastruct.Testimonials
	var users []datastruct.Users

	tx := t.db.Preload("UserTestimonials").Find(&testimonials)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to get testimonials: %v", tx.Error)
	}

	data := make([]dto.Testimonials, 0)
	for _, testimonial := range testimonials {
		tx := t.db.Preload("UserDetails").
			Where("users.id = ?", testimonial.UserTestimonials.UserID).
			First(&users)
		if tx.Error != nil {
			return nil, fmt.Errorf("failed to get user: %v", tx.Error)
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
		data = append(data, newTestimonials)
	}

	return data, nil
}

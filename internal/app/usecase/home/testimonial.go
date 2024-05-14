package home

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) Testimonial() ([]*response.TestimonialHome, error) {

	testimonials, testimonialMediaIDs, err := u.localHomeRepo.FindTestimonialInfo()
	if err != nil {
		return nil, err
	}

	mediaMaps, err := u.localHomeRepo.FindMediaInfo(testimonialMediaIDs)
	if err != nil {
		return nil, err
	}

	mediaMapTestimonial, mediaMentorFound := mediaMaps["media1"]
	if !mediaMentorFound {
		return nil, fmt.Errorf("failed to find  media mentor")
	}

	results := make([]*response.TestimonialHome, 0)
	for _, testimonial := range testimonials {
		results = append(results, &response.TestimonialHome{
			UUID:    testimonial.UUID,
			Name:    testimonial.Name,
			Class:   testimonial.Class,
			Media:   helper.MultiResImages(mediaMapTestimonial[testimonial.MediaID]),
			Comment: testimonial.Comment,
			Rating:  testimonial.Rating,
		})
	}

	if len(results) == 0 {
		return nil, helper.ErrDataNotFound
	}

	return results, nil
}

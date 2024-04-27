package home

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Testimonial() ([]*response.TestimonialHome, error) {
	userTestimonials, err := u.userTestimonialRepo.FindBy(map[string]interface{}{}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find user testimonials: %v", err)
	}

	userTestimonialUserIDs := make([]int, len(userTestimonials))
	for i, userTestimonial := range userTestimonials {
		userTestimonialUserIDs[i] = userTestimonial.UserID
	}

	user, err := u.userRepo.FindBy(map[string]interface{}{
		"id": userTestimonialUserIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %v", err)
	}

	userIDs := make([]int, len(user))
	for i, u := range user {
		userIDs[i] = u.ID
	}

	userDetails, err := u.userDetailRepo.FindBy(map[string]interface{}{
		"user_id": userIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find user details: %v", err)
	}

	userDetailMap := make(map[int]*model.UserDetail)
	for _, userDetail := range userDetails {
		userDetailMap[userDetail.UserID] = userDetail
	}

	classUserIDs := make([]int, len(user))
	for i, ud := range userDetails {
		classUserIDs[i] = ud.ClassUserID
	}

	classes, err := u.classUserRepo.FindBy(map[string]interface{}{
		"id": classUserIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find classes: %v", err)
	}

	classMap := make(map[int]*model.ClassUser)
	for _, class := range classes {
		classMap[class.ID] = class
	}

	results := make([]*response.TestimonialHome, len(userTestimonials))
	for i, userTestimonial := range userTestimonials {
		user, ok := userDetailMap[userTestimonial.UserID]
		if !ok {
			continue
		}

		results[i] = &response.TestimonialHome{
			UUID:           user.UUID,
			Name:           user.Name,
			Class:          classMap[user.ClassUserID].Name,
			ProfilePicture: user.ProfilePicture,
			Comment:        userTestimonial.Comment,
			Rating:         userTestimonial.Rating,
		}
	}

	if len(results) == 0 {
		return nil, helper.ErrDataNotFound
	}

	return results, nil
}

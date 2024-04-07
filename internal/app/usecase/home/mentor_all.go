package home

import (
	"fmt"
	"math"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) MentorAll() ([]*response.MentorHome, error) {
	role, err := u.roleRepo.FindOneBy(map[string]interface{}{
		"slug": "staff-bimbel",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find role: %v", err)
	}

	users, err := u.userRepo.FindBy(map[string]interface{}{
		"role_id": role.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find users: %v", err)
	}

	userIDs := make([]int, len(users))
	for i, user := range users {
		userIDs[i] = user.ID
	}

	userDetails, err := u.userDetailRepo.FindBy(map[string]interface{}{
		"user_id": userIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find user details: %v", err)
	}

	userDetailMap := make(map[int]*model.UserDetail)
	for _, userDetail := range userDetails {
		userDetailMap[userDetail.ID] = userDetail
	}

	mentors, err := u.mentorRepo.FindBy(map[string]interface{}{
		"user_id": userIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find mentors: %v", err)
	}

	mentorIDs := make([]int, len(mentors))
	for i, mentor := range mentors {
		mentorIDs[i] = mentor.ID
	}

	modules, err := u.moduleRepo.FindBy(map[string]interface{}{
		"id": mentorIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find modules: %v", err)
	}

	moduleMap := make(map[int]*model.Module)
	for _, module := range modules {
		moduleMap[module.ID] = module
	}

	userMentorTestimonials, err := u.userMentorTestimonialRepo.FindBy(map[string]interface{}{
		"mentor_id": mentorIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find user mentor testimonials: %v", err)
	}

	mentorStats := make(map[int]response.MentorStats)
	for _, umt := range userMentorTestimonials {
		stats := mentorStats[umt.MentorID]
		stats.Rating = (stats.Rating*float64(stats.Testimonial) + float64(umt.Rating)) / float64(stats.Testimonial+1)
		stats.Testimonial++
		mentorStats[umt.MentorID] = stats
	}

	results := make([]*response.MentorHome, len(mentors))
	for i, mentor := range mentors {
		moduleName := moduleMap[mentor.ModuleID].Name
		userName := userDetailMap[mentor.UserID].Name
		profilePicture := userDetailMap[mentor.UserID].ProfilePicture
		rating := math.Round(float64(mentorStats[mentor.ID].Rating)*10) / 10

		results[i] = &response.MentorHome{
			UUID:           mentor.UUID,
			Description:    mentor.Description,
			Motto:          mentor.Motto,
			TeachingField:  moduleName,
			Name:           userName,
			ProfilePicture: profilePicture,
			Rating:         rating,
		}
	}

	if len(results) == 0 {
		return nil, helper.ErrDataNotFound
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Rating > results[j].Rating
	})

	return results, nil
}

package home

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) MentorAll() ([]*response.MentorHome, error) {
	role, err := u.roleRepo.FindOneBy(map[string]interface{}{
		"slug": "mentor",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to finding role: %v", err)
	}

	users, err := u.userRepo.FindBy(map[string]interface{}{
		"role_id": role.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to finding user: %v", err)
	}

	userIDs := make([]int, 0)
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}

	userDetails, err := u.userDetailRepo.FindBy(map[string]interface{}{
		"user_id": userIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to finding user detail: %v", err)
	}

	userDetailNameMap := make(map[int]string)
	for _, userDetail := range userDetails {
		userDetailNameMap[userDetail.ID] = userDetail.Name
	}

	userDetailProfilePhotoMap := make(map[int]string)
	for _, userDetail := range userDetails {
		userDetailProfilePhotoMap[userDetail.ID] = userDetail.ProfilePicture
	}

	mentors, err := u.mentorRepo.FindBy(map[string]interface{}{
		"user_id": userIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to finding mentors: %v", err)
	}

	mentorMap := make(map[int]*model.Mentor)
	for _, mentor := range mentors {
		mentorMap[mentor.ID] = mentor
	}

	results := make([]*response.MentorHome, 0)
	for _, mentor := range mentors {
		result := &response.MentorHome{
			UUID:           mentor.UUID,
			Type:           mentor.Type,
			Description:    mentor.Description,
			Motto:          mentor.Motto,
			Name:           userDetailNameMap[mentor.UserID],
			ProfilePicture: userDetailProfilePhotoMap[mentor.UserID],
		}

		results = append(results, result)
	}

	if len(results) == 0 {
		return nil, helper.ErrDataNotFound
	}

	return results, nil
}

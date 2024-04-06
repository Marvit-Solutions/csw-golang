package home

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) TopMentor() ([]*response.MentorHome, error) {
	role, err := u.roleRepo.FindOneBy(map[string]interface{}{
		"slug": "mentor",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to finding role: %v", err)
	}

	users, err := u.userRepo.FindBy(map[string]interface{}{
		"role_id": role.ID,
	}, 1, 3)
	if err != nil {
		return nil, fmt.Errorf("failed to finding user: %v", err)
	}

	userIDs := make([]int, 0)
	for _, user := range users {
		userIDs = append(userIDs, user.ID)
	}

	userDetails, err := u.userDetailRepo.FindBy(map[string]interface{}{
		"user_id": userIDs,
	}, 1, 3)
	if err != nil {
		return nil, fmt.Errorf("failed to finding user detail: %v", err)
	}

	userDetailMap := make(map[int]*model.UserDetail)
	for _, userDetail := range userDetails {
		userDetailMap[userDetail.ID] = userDetail
	}

	mentors, err := u.mentorRepo.FindBy(map[string]interface{}{
		"user_id": userIDs,
	}, 1, 3)
	if err != nil {
		return nil, fmt.Errorf("failed to finding mentors: %v", err)
	}

	mentorMap := make(map[int]*model.Mentor)
	for _, mentor := range mentors {
		mentorMap[mentor.ID] = mentor
	}

	results := make([]*response.MentorHome, 0)
	for _, mentor := range mentors {
		switch mentor.Type {
		case "TIU", "TKP", "Matematika":
			result := &response.MentorHome{
				UUID:           mentor.UUID,
				Type:           mentor.Type,
				Description:    mentor.Description,
				Motto:          mentor.Motto,
				Name:           userDetailMap[mentor.UserID].Name,
				ProfilePicture: userDetailMap[mentor.UserID].ProfilePicture,
			}
			results = append(results, result)
		}
	}

	return results, nil
}

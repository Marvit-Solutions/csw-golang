package home

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) TopMentor() (*response.AllMentor, error) {
	role, err := u.roleRepo.FindOneBy(map[string]interface{}{
		"name": "mentor",
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

	userMap := make(map[int]*model.User)
	for _, user := range users {
		userMap[user.ID] = &model.User{
			ID: user.ID,
		}
	}

	var mentorIDs []int
	for userID := range userMap {
		mentorIDs = append(mentorIDs, userID)
	}

	mentors, err := u.mentorRepo.FindBy(map[string]interface{}{
		"user_id": mentorIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to finding mentors: %v", err)
	}

	mentorMap := make(map[int]*model.Mentor)
	for _, mentor := range mentors {
		mentorMap[mentor.ID] = mentor
	}

	var result []response.TopMentor

	return nil, nil
}

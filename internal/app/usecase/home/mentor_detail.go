package home

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) MentorDetail(req request.ParamMentorDetailHome) (*response.MentorDetailHome, error) {
	mentor, err := u.mentorRepo.FindOneBy(map[string]interface{}{
		"uuid": req.UUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find mentor: %v", err)
	}

	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"id": mentor.UserID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	userDetail, err := u.userDetailRepo.FindOneBy(map[string]interface{}{
		"user_id": user.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find user detail: %v", err)
	}

	module, err := u.moduleRepo.FindOneBy(map[string]interface{}{
		"id": mentor.ModuleID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find module: %v", err)
	}

	uniques, err := u.uniqueRepo.FindBy(map[string]interface{}{
		"mentor_id": mentor.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find uniques: %v", err)
	}

	uniqueMap := make(map[int]*model.Unique)
	for _, unique := range uniques {
		uniqueMap[unique.ID] = unique
	}

	result := &response.MentorDetailHome{
		UUID:           mentor.UUID,
		Name:           userDetail.Name,
		TeachingField:  module.Name,
		Description:    mentor.Description,
		ProfilePicture: userDetail.ProfilePicture,
	}

	for _, unique := range uniqueMap {
		result.Uniques = append(result.Uniques, uniqueMap[unique.ID].Uniqueness)
	}

	return result, nil
}

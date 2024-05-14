package home

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) MentorDetail(req request.ParamMentorDetailHome) (*response.MentorDetailHome, error) {
	mentor, err := u.localHomeRepo.FindMentorDetailInfo(req)
	if err != nil {
		return nil, err
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

	var media *model.Media
	if mentor.MediaID != 0 {
		media, err = u.mediaRepo.FindOneBy(map[string]interface{}{
			"id": mentor.MediaID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to find media: %v", err)
		}
	}

	result := &response.MentorDetailHome{
		UUID:          mentor.UUID,
		Name:          mentor.Name,
		TeachingField: mentor.TeachingField,
		Description:   mentor.Description,
		Media:         helper.MultiResImages(media),
	}

	for _, unique := range uniqueMap {
		result.Uniques = append(result.Uniques, uniqueMap[unique.ID].Uniqueness)
	}

	if result.UUID == "" {
		return nil, helper.ErrDataNotFound
	}

	return result, nil
}

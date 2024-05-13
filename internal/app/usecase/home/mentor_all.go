package home

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) MentorAll() ([]*response.MentorHome, error) {

	mentors, mentorMediaIDs, err := u.localHomeRepo.FindMentorInfo()
	if err != nil {
		return nil, err
	}

	mediaMaps, err := u.localHomeRepo.FindMediaInfo(mentorMediaIDs)
	if err != nil {
		return nil, err
	}

	mediaMapMentor, mediaMentorFound := mediaMaps["media1"]
	if !mediaMentorFound {
		return nil, fmt.Errorf("failed to find  media mentor")
	}

	results := make([]*response.MentorHome, 0)
	for _, mentor := range mentors {
		results = append(results, &response.MentorHome{
			UUID:          mentor.UUID,
			Name:          mentor.Name,
			TeachingField: mentor.TeachingField,
			Description:   mentor.Description,
			Motto:         mentor.Motto,
			Media:         helper.MultiResImages(mediaMapMentor[mentor.MediaID]),
			Rating:        mentor.Rating,
		})
	}

	if len(results) == 0 {
		return nil, helper.ErrDataNotFound
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Rating > results[j].Rating
	})

	return results, nil
}

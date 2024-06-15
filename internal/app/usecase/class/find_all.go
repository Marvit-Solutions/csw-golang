package class

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) FindAll(req request.Class) ([]*response.Class, error) {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"id":      req.AuthenticatedUser,
		"role_id": helper.PembeliPaketBimbel,
	})
	if user == nil {
		return nil, helper.ErrAccessDenied
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	schedules, scheduleMediaIDs, mentorMediaIDs, err := u.classLocalRepo.FindScheduleInfo(req)
	if err != nil {
		return nil, err
	}

	mediaMaps, err := u.classLocalRepo.FindMediaInfo(scheduleMediaIDs, mentorMediaIDs)
	if err != nil {
		return nil, err
	}

	mediaMapSchedule, scheduleFound := mediaMaps["media1"]
	mediaMapMentor, mentorFound := mediaMaps["media2"]

	if !scheduleFound || !mentorFound {
		return nil, fmt.Errorf("failed to find schedule and mentor media")
	}

	result := make([]*response.Class, 0)
	for _, schedule := range schedules {
		var mentorMedia *model.MultiResImage
		if schedule.MentorMediaID != nil {
			mentorMedia = helper.MultiResImages(mediaMapMentor[*schedule.MentorMediaID])
		}

		var scheduleMedia *model.MultiResImage
		if schedule.ScheduleMediaID != nil {
			scheduleMedia = helper.MultiResImages(mediaMapSchedule[*schedule.ScheduleMediaID])
		}

		result = append(result, &response.Class{
			UUID:        schedule.UUID,
			Module:      schedule.ModuleName,
			MeetingDate: schedule.MeetingDate,
			Media:       scheduleMedia,
			Subject: response.SubjectClass{
				UUID:        schedule.SubjectUUID,
				Name:        schedule.SubjectName,
				Description: schedule.SubjectDescription,
			},
			Mentor: response.MentorClass{
				UUID:  schedule.MentorUUID,
				Name:  schedule.ShortName,
				Media: mentorMedia,
			},
			TotalParticipant: 0,
			Duration:         schedule.Duration,
		})
	}

	return result, nil
}

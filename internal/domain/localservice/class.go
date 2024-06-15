package localservice

import (
	"fmt"
	"strings"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localrepository"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"gorm.io/gorm"
)

type ClassService struct {
	DB *gorm.DB
}

func NewClassService(
	DB *gorm.DB,
) localrepository.Class {
	return &ClassService{
		DB,
	}
}

func (svc *ClassService) FindScheduleInfo(req request.Class) ([]*response.ClassScheduleList, []int, []int, error) {
	schedules := make([]*response.ClassScheduleList, 0)

	res := svc.DB.
		Select(`s.uuid,s.media_id AS "schedule_media_id", cpt.id AS "class_plan_type_id", cpt.name AS "class_name", m.uuid AS "mentor_uuid",m.short_name,
				ud.media_id AS "mentor_media_id", sj.uuid AS "subject_uuid",sj.name AS "subject_name", sj.description AS "subject_description", s.meeting_date,
				s.meeting_link, s.duration,mdl.name AS "module_name", mdl.slug AS "module_slug"`).
		Table("Schedules s").
		Joins(`LEFT JOIN subjects sj ON sj.id = s.subject_id`).
		Joins(`LEFT JOIN sub_modules sm ON sm.id = sj.sub_module_id`).
		Joins(`LEFT JOIN modules mdl ON mdl.id = sm.module_id`).
		Joins(`LEFT JOIN class_plan_types cpt ON cpt.id = s.class_plan_type_id`).
		Joins(`LEFT JOIN mentors m ON m.id = s.mentor_id`).
		Joins(`LEFT JOIN users u ON u.id = m.user_id`).
		Joins(`LEFT JOIN user_details ud ON ud.user_id = u.id`).
		Where(`mdl.slug = ?`, strings.ToLower(req.Module))

	err := res.Scan(&schedules).Error
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to find schedules: %v", err)
	}

	var scheduleMediaIDs, mentorMediaIDs []int
	for _, schedule := range schedules {
		if schedule.ScheduleMediaID != nil {
			scheduleMediaIDs = append(scheduleMediaIDs, *schedule.ScheduleMediaID)
		}
		if schedule.MentorMediaID != nil {
			mentorMediaIDs = append(mentorMediaIDs, *schedule.MentorMediaID)
		}
	}

	return schedules, scheduleMediaIDs, mentorMediaIDs, nil
}

func (svc *ClassService) FindMediaInfo(mediaIDs ...[]int) (map[string]map[int]*model.Media, error) {
	allMediaIDs := make([]int, 0)
	for _, ids := range mediaIDs {
		allMediaIDs = append(allMediaIDs, ids...)
	}

	medias := make([]*model.Media, 0)
	err := svc.DB.Where("id IN ?", allMediaIDs).Find(&medias).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch media: %v", err)
	}

	mediaMap := make(map[int]*model.Media)
	for _, media := range medias {
		mediaMap[media.ID] = media
	}

	result := make(map[string]map[int]*model.Media)
	for idx, ids := range mediaIDs {
		mediaMapByType := make(map[int]*model.Media)
		for _, id := range ids {
			if media, ok := mediaMap[id]; ok {
				mediaMapByType[id] = media
			}
		}
		result[fmt.Sprintf("media%d", idx+1)] = mediaMapByType
	}

	return result, nil
}

package localservice

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localrepository"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"gorm.io/gorm"
)

type HomeService struct {
	DB *gorm.DB
}

func NewHomeService(
	DB *gorm.DB,
) localrepository.Home {
	return &HomeService{
		DB,
	}
}

func (svc *HomeService) FindMentorInfo(orderByRating bool) ([]*response.MentorHomeList, []int, error) {
	mentors := make([]*response.MentorHomeList, 0)

	var orderByClause string
	if orderByRating {
		orderByClause = "ORDER BY rating DESC LIMIT 3"
	}

	query := fmt.Sprintf(`SELECT m.uuid, ud.media_id, m.short_name AS name, mdl.name AS teaching_field, m.description, m.motto, COALESCE(ROUND(avg_rating.avg_rating, 2), 0) AS rating
			FROM mentors m
			LEFT JOIN user_details ud ON ud.user_id = m.user_id
			LEFT JOIN modules mdl ON mdl.id = m.module_id
			LEFT JOIN (
				SELECT mentor_id, AVG(rating) AS avg_rating 
				FROM user_mentor_testimonials 
				GROUP BY 
				mentor_id
				) AS avg_rating ON avg_rating.mentor_id = m.id %s`, orderByClause)

	res := svc.DB.Raw(query)

	err := res.Scan(&mentors).Error
	if err != nil {
		return nil, nil, fmt.Errorf("failed to find mentors: %v", err)
	}

	var mentorMediaIDs []int
	for _, mentor := range mentors {
		mentorMediaIDs = append(mentorMediaIDs, mentor.MediaID)
	}

	return mentors, mentorMediaIDs, nil
}

func (svc *HomeService) FindMentorDetailInfo(req request.ParamMentorDetailHome) (*response.MentorDetailHomeList, error) {
	voucher := &response.MentorDetailHomeList{}

	res := svc.DB.Select(`m.id ,m.uuid, ud.media_id, m.short_name AS name, mdl.name AS teaching_field, m.description`).
		Table(`mentors m`).
		Joins(`LEFT JOIN user_details ud ON ud.user_id = m.user_id`).
		Joins(`LEFT JOIN modules mdl ON mdl.id = m.module_id`).
		Where(`m.uuid = ?`, req.UUID)

	err := res.Scan(voucher).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find mentor: %v", err)
	}

	return voucher, nil
}

func (svc *HomeService) FindMediaInfo(mediaIDs ...[]int) (map[string]map[int]*model.Media, error) {
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

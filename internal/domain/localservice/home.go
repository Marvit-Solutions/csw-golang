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

func (svc *HomeService) FindPlanInfo(req request.PlanHome, orderByPrice bool) ([]*response.PlanHomeList, []int, error) {
	plans := make([]*response.PlanHomeList, 0)

	res := svc.DB.Select(`p.uuid, p.media_id, m.name AS module_name, p.name, p.price, p.group, p.exercise, p.access, p.module, p.try_out, p.zoom`).
		Table(`plans p`).
		Joins(`LEFT JOIN modules m ON m.id = p.module_id`)

	if orderByPrice {
		res = res.Order(`p.price ASC LIMIT 3`)
	} else {
		if req.Module != "" {
			module := strings.Split(strings.ToLower(req.Module), ",")
			res = res.Where("m.slug LIKE ?", "%"+module[0]+"%")
		}
		if req.Name != "" {
			planName := strings.Split(strings.ToLower(req.Name), ",")
			res = res.Where("(p.slug LIKE ? OR p.name LIKE ?)", "%"+planName[0]+"%", "%"+planName[0]+"%")
		}
	}

	err := res.Scan(&plans).Error
	if err != nil {
		return nil, nil, fmt.Errorf("failed to find plan: %v", err)
	}

	var planMediaIDs []int
	for _, plan := range plans {
		planMediaIDs = append(planMediaIDs, plan.MediaID)
	}

	return plans, planMediaIDs, nil
}

func (svc *HomeService) FindTestimonialInfo() ([]*response.TestimonialHomeList, []int, error) {
	testimonial := make([]*response.TestimonialHomeList, 0)

	res := svc.DB.Select(`ut.uuid, ud.media_id, ud.name, cu.name AS "class", ut.comment, ut.rating`).
		Table(`user_testimonials ut`).
		Joins(`LEFT JOIN user_details ud ON ud.user_id = ut.user_id`).
		Joins(`LEFT JOIN class_users cu ON cu.id = ud.class_user_id`)

	err := res.Scan(&testimonial).Error
	if err != nil {
		return nil, nil, fmt.Errorf("failed to find testimonial: %v", err)
	}

	var testimonialMediaIDs []int
	for _, testimonial := range testimonial {
		testimonialMediaIDs = append(testimonialMediaIDs, testimonial.MediaID)
	}

	return testimonial, testimonialMediaIDs, nil
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

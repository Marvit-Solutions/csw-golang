package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
)

type Home interface {
	FindMentorInfo(orderByRating bool) ([]*response.MentorHomeList, []int, error)
	FindMentorDetailInfo(req request.ParamMentorDetailHome) (*response.MentorDetailHomeList, error)
	FindPlanInfo(req request.PlanHome, orderByPrice bool) ([]*response.PlanHomeList, []int, error)
	FindTestimonialInfo() ([]*response.TestimonialHomeList, []int, error)
	FindMediaInfo(mediaIDs ...[]int) (map[string]map[int]*model.Media, error)
}

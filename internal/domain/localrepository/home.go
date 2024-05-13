package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

type Home interface {
	FindMentorInfo(orderByRating bool) ([]*response.MentorHomeList, []int, error)
	FindMentorDetailInfo(req request.ParamMentorDetailHome) (*response.MentorDetailHomeList, error)
	FindPlanInfo(req request.PlanHome) ([]*response.PlanHomeList, []int, error)
	FindMediaInfo(mediaIDs ...[]int) (map[string]map[int]*model.Media, error)
}

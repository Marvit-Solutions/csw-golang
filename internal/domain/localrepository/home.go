package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

type Home interface {
	FindMentorInfo() ([]*response.MentorHomeList, []int, error)
	FindMediaInfo(mediaIDs ...[]int) (map[string]map[int]*model.Media, error)
}

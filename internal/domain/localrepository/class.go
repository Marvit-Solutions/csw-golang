package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

type Class interface {
	FindScheduleInfo(req request.Class) ([]*response.ClassScheduleList, []int, []int, error)
	FindMediaInfo(mediaIDs ...[]int) (map[string]map[int]*model.Media, error)
}

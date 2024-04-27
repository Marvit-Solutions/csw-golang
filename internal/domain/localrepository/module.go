package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

type Module interface {
	FindSubjectInfo(result *response.MaterialFindResponse) ([]*model.Subject, []*model.SubModule, []*model.Module, map[int]int, error)
	MapSubjectInfo(subjects []*model.Subject, subModules []*model.SubModule, modules []*model.Module, countBySubjectID map[int]int) ([]*response.SubjectFindResponse, error)
	FindSubSubjectInfo(result *response.MaterialFindResponse) ([]*model.Subject, []*model.SubSubject, []*model.SubModule, []*model.Module, error)
	MapSubSubjectInfo(subjects []*model.Subject, subSubjects []*model.SubSubject, subModules []*model.SubModule, modules []*model.Module) ([]*response.SubSubjectFindResponse, error)
	// SortSubjectInfo(result *response.MaterialFindResponse)
	// SortSubSubjectInfo(result *response.MaterialFindResponse)
}

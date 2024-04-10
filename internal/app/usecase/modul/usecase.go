package modul

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	ModulAll() ([]*response.ModulResponse, error)
	ModulDetail(request.ParamModulDetail) (*response.ModulDetailResponse, error)
}

type usecase struct {
	db            *gorm.DB
	subModuleRepo repository.SubModuleRepository
	moduleRepo    repository.ModuleRepository
	subjectRepo   repository.SubjectRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:            db,
		moduleRepo:    service.NewModuleService(db, nil),
		subModuleRepo: service.NewSubModuleService(db, nil),
		subjectRepo:   service.NewSubjectService(db, nil),
	}
}

package module

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localrepository"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localservice"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	ModuleAll() ([]*response.ModulResponse, error)
	ModuleDetail(req request.ParamModule) (*response.ModuleDetailResponse, error)
	MaterialAll(req request.ParamModule) (*response.MaterialResponse, error)
	MaterialFind(req request.Material) (*response.MaterialFindResponse, error)
}

type usecase struct {
	db              *gorm.DB
	subModuleRepo   repository.SubModuleRepository
	moduleRepo      repository.ModuleRepository
	subjectRepo     repository.SubjectRepository
	subSubjectRepo  repository.SubSubjectRepository
	quizRepo        repository.QuizRepository
	testTypeRepo    repository.TestTypeRepository
	moduleLocalRepo localrepository.Module
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db: db,
		moduleLocalRepo: localservice.NewModuleService(
			db,
			service.NewSubjectService(db, nil),
			service.NewSubSubjectService(db, nil),
			service.NewSubModuleService(db, nil),
			service.NewModuleService(db, nil),
		),
		moduleRepo:     service.NewModuleService(db, nil),
		subModuleRepo:  service.NewSubModuleService(db, nil),
		subjectRepo:    service.NewSubjectService(db, nil),
		subSubjectRepo: service.NewSubSubjectService(db, nil),
		quizRepo:       service.NewQuizService(db, nil),
		testTypeRepo:   service.NewTestTypeService(db, nil),
	}
}

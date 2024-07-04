package dashboard

import (
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localrepository"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localservice"
	"gorm.io/gorm"
)

type Usecase interface {
	FindMaterial(req request.ParamDashboard) ([]*response.MaterialDashboard, error)
	QuizAll(req request.ParamDashboard) ([]*response.QuizAllDashboardPerTest, error)
}

type usecase struct {
	db                 *gorm.DB
	userRepo           repository.UserRepository
	subjectRepo        repository.SubjectRepository
	moduleRepo         repository.ModuleRepository
	subModuleRepo      repository.SubModuleRepository
	testTypeRepo       repository.TestTypeRepository
	dashboardLocalRepo localrepository.Dashboard
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:                 db,
		userRepo:           service.NewUserService(db, nil),
		subjectRepo:        service.NewSubjectService(db, nil),
		moduleRepo:         service.NewModuleService(db, nil),
		subModuleRepo:      service.NewSubModuleService(db, nil),
		testTypeRepo:       service.NewTestTypeService(db, nil),
		dashboardLocalRepo: localservice.NewDashboardService(db),
	}
}

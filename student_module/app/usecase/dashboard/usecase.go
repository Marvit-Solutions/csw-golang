package dashboard

import (
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
	"gorm.io/gorm"
)

type Usecase interface {
	FindMaterial(req request.Dashboard) ([]*response.MaterialDashboard, error)
}

type usecase struct {
	db            *gorm.DB
	userRepo      repository.UserRepository
	subjectRepo   repository.SubjectRepository
	moduleRepo    repository.ModuleRepository
	subModuleRepo repository.SubModuleRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:            db,
		userRepo:      service.NewUserService(db, nil),
		subjectRepo:   service.NewSubjectService(db, nil),
		moduleRepo:    service.NewModuleService(db, nil),
		subModuleRepo: service.NewSubModuleService(db, nil),
	}
}

package home

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	MentorTop() ([]*response.MentorHome, error)
	MentorAll() ([]*response.MentorHome, error)
	MentorDetail(request.ParamMentorDetailHome) (*response.MentorDetailHome, error)
	// TopPlan() ([]*response.PlanHome, error)
	// AllPlan() ([]*response.PlanHome, error)
}

type usecase struct {
	db                        *gorm.DB
	userRepo                  repository.UserRepository
	userDetailRepo            repository.UserDetailRepository
	userMentorTestimonialRepo repository.UserMentorTestimonialRepository
	roleRepo                  repository.RoleRepository
	mentorRepo                repository.MentorRepository
	moduleRepo                repository.ModuleRepository
	uniqueRepo                repository.UniqueRepository
}

func NewUsecase(
	db *gorm.DB,
	conf config.Config,
) Usecase {
	return &usecase{
		db:                        db,
		userRepo:                  service.NewUserService(db, nil),
		userDetailRepo:            service.NewUserDetailService(db, nil),
		userMentorTestimonialRepo: service.NewUserMentorTestimonialService(db, nil),
		roleRepo:                  service.NewRoleService(db, nil),
		mentorRepo:                service.NewMentorService(db, nil),
		moduleRepo:                service.NewModuleService(db, nil),
		uniqueRepo:                service.NewUniqueService(db, nil),
	}
}

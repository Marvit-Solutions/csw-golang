package home

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	MentorTop() ([]*response.MentorHome, error)
	MentorAll() ([]*response.MentorHome, error)
	MentorDetail(request.ParamMentorDetailHome) (*response.MentorDetailHome, error)
	PlanTop() ([]*response.PlanHome, error)
	PlanAll(req request.PlanHome) ([]*response.PlanHome, error)
	Testimonial() ([]*response.TestimonialHome, error)
}

type usecase struct {
	db                        *gorm.DB
	userRepo                  repository.UserRepository
	userDetailRepo            repository.UserDetailRepository
	userMentorTestimonialRepo repository.UserMentorTestimonialRepository
	userTestimonialRepo       repository.UserTestimonialRepository
	classUserRepo             repository.ClassUserRepository
	roleRepo                  repository.RoleRepository
	mentorRepo                repository.MentorRepository
	moduleRepo                repository.ModuleRepository
	uniqueRepo                repository.UniqueRepository
	planRepo                  repository.PlanRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:                        db,
		userRepo:                  service.NewUserService(db, nil),
		userDetailRepo:            service.NewUserDetailService(db, nil),
		userMentorTestimonialRepo: service.NewUserMentorTestimonialService(db, nil),
		userTestimonialRepo:       service.NewUserTestimonialService(db, nil),
		classUserRepo:             service.NewClassUserService(db, nil),
		roleRepo:                  service.NewRoleService(db, nil),
		mentorRepo:                service.NewMentorService(db, nil),
		moduleRepo:                service.NewModuleService(db, nil),
		uniqueRepo:                service.NewUniqueService(db, nil),
		planRepo:                  service.NewPlanService(db, nil),
	}
}

package home

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
	MentorTop() ([]*response.MentorHome, error)
	MentorAll() ([]*response.MentorHome, error)
	MentorDetail(request.ParamMentorDetailHome) (*response.MentorDetailHome, error)
	PlanTop() ([]*response.PlanHome, error)
	PlanAll(req request.PlanHome) ([]*response.PlanHome, error)
	Testimonial() ([]*response.TestimonialHome, error)
}

type usecase struct {
	db            *gorm.DB
	localHomeRepo localrepository.Home
	uniqueRepo    repository.UniqueRepository
	mediaRepo     repository.MediaRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:            db,
		localHomeRepo: localservice.NewHomeService(db),
		uniqueRepo:    service.NewUniqueService(db, nil),
		mediaRepo:     service.NewMediaService(db, nil),
	}
}

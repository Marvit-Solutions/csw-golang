package home

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	TopMentor() (*response.AllMentor, error)
	AllMentor() (*response.AllMentor, error)
}

type usecase struct {
	db         *gorm.DB
	userRepo   repository.UserRepository
	roleRepo   repository.RoleRepository
	mentorRepo repository.MentorRepository
}

func NewUsecase(
	db *gorm.DB,
	conf config.Config,
) Usecase {
	return &usecase{
		db:         db,
		userRepo:   service.NewUserService(db, nil),
		roleRepo:   service.NewRoleService(db, nil),
		mentorRepo: service.NewMentorService(db, nil),
	}
}

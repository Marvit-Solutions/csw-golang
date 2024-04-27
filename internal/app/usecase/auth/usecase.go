package auth

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"

	"gorm.io/gorm"
)

type Usecase interface {
	Register(req request.RegisterRequest) (*response.AuthResponse, error)
	Login(req request.LoginRequest) (*response.AuthResponse, error)
}

type usecase struct {
	db             *gorm.DB
	userRepo       repository.UserRepository
	classUserRepo  repository.ClassUserRepository
	userDetailRepo repository.UserDetailRepository
	roleRepo       repository.RoleRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:             db,
		userRepo:       service.NewUserService(db, nil),
		classUserRepo:  service.NewClassUserService(db, nil),
		userDetailRepo: service.NewUserDetailService(db, nil),
		roleRepo:       service.NewRoleService(db, nil),
	}
}

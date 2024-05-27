package user

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	Find(req request.User) (*response.User, error)
}

type usecase struct {
	db             *gorm.DB
	userRepo       repository.UserRepository
	userDetailRepo repository.UserDetailRepository
	classUserRepo  repository.ClassUserRepository
	mediaRepo      repository.MediaRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:             db,
		userRepo:       service.NewUserService(db, nil),
		userDetailRepo: service.NewUserDetailService(db, nil),
		classUserRepo:  service.NewClassUserService(db, nil),
		mediaRepo:      service.NewMediaService(db, nil),
	}
}

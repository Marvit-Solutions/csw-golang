package class

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
	FindAll(req request.Class) ([]*response.Class, error)
}

type usecase struct {
	db                *gorm.DB
	userRepo          repository.UserRepository
	classLocalRepo    localrepository.Class
	classPlanUserRepo repository.ClassPlanUserRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:                db,
		classLocalRepo:    localservice.NewClassService(db),
		userRepo:          service.NewUserService(db, nil),
		classPlanUserRepo: service.NewClassPlanUserService(db, nil),
	}
}

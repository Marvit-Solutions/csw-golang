package location

import (
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"

	"gorm.io/gorm"
)

type Usecase interface {
	Province(req request.LocationRequest) ([]*response.LocationResponse, error)
	Regency(req request.LocationRequest) ([]*response.LocationResponse, error)
	District(req request.LocationRequest) ([]*response.LocationResponse, error)
}

type usecase struct {
	db *gorm.DB
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db: db,
	}
}

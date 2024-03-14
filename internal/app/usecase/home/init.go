package home

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"gorm.io/gorm"
)

type Usecase interface {
	TopMentor() (*response.Mentor, error)
	AllMentor() (*response.Mentor, error)
}

type usecase struct {
	db *gorm.DB
}

func NewUsecase(
	db *gorm.DB,
	conf config.Config,
) Usecase {
	return &usecase{
		db: db,
	}
}

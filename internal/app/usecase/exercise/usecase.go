package exercise

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"gorm.io/gorm"
)

type Usecase interface {
	FindAll(req request.ParamExercise) ([]*response.ExerciseResponse, error)
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

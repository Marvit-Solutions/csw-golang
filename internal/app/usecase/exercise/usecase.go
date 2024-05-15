package exercise

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/service"
	"gorm.io/gorm"
)

type Usecase interface {
	FindAll(req request.ParamExercise) ([]*response.ExerciseResponse, error)
}

type usecase struct {
	db           *gorm.DB
	testTypeRepo repository.TestTypeRepository
	moduleRepo   repository.ModuleRepository
	exerciseRepo repository.ExerciseRepository
}

func NewUsecase(
	db *gorm.DB,
) Usecase {
	return &usecase{
		db:           db,
		testTypeRepo: service.NewTestTypeService(db, nil),
		moduleRepo:   service.NewModuleService(db, nil),
		exerciseRepo: service.NewExerciseService(db, nil),
	}
}

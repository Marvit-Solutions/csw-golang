package test

import (
	"csw-golang/internal/domain/entity/dto"
	et "csw-golang/internal/domain/repository/exercise/test"
)

type ExerciseTestUsecase interface {
	AddExerciseTestUsecase(request dto.SubjectTestTypeExercisesRequest) (dto.SubjectTestTypeExercisesResponse, error)
	GetExerciseTestUsecase(id string) (dto.SubjectTestTypeExercisesResponse, error)
	GetAllExerciseTestUsecase() ([]dto.SubjectTestTypeExercisesResponse, error)
}

type exerciseTestUsecase struct {
	exerciseTestRepo et.ExerciseTestRepo
}

func New(exerciseTestRepo et.ExerciseTestRepo) ExerciseTestUsecase {
	return &exerciseTestUsecase{
		exerciseTestRepo,
	}
}

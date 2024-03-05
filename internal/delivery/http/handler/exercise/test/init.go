package test

import (
	mt "csw-golang/internal/domain/usecase/exercise/test"
)

type ExerciseTestHandler struct {
	exerciseTestUsecase mt.ExerciseTestUsecase
}

func New(exerciseTestUsecase mt.ExerciseTestUsecase) *ExerciseTestHandler {
	return &ExerciseTestHandler{
		exerciseTestUsecase,
	}
}

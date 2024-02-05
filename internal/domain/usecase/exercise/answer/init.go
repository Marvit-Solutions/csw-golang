package answer

import (
	"csw-golang/internal/domain/entity/dto"
	es "csw-golang/internal/domain/repository/exercise/answer"
)

type ExerciseAnswerUsecase interface {
	AddAnswerUsecase(answer dto.UserSubmittedAnswerExercisesRequest) (dto.UserSubmittedAnswerExercisesResponse, error)
}

type exerciseAnswerUsecase struct {
	exerciseAnswerRepo es.ExerciseAnswerRepo
}

func New(exerciseAnswerRepo es.ExerciseAnswerRepo) ExerciseAnswerUsecase {
	return &exerciseAnswerUsecase{
		exerciseAnswerRepo,
	}
}

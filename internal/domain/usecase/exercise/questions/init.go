package questions

import (
	"csw-golang/internal/domain/entity/dto"
	eq "csw-golang/internal/domain/repository/exercise/question"
)

type ExerciseQuestionsUsecase interface {
	AddExerciseQuestion(exerciseQuestion dto.QuestionExercisesRequest) (dto.QuestionExercisesResponse, error)
	GetExerciseQuestions(exerciseQuestionsID string) (dto.QuestionExercisesResponse, error)
}

type exerciseQuestionsUsecase struct {
	exerciseQuestionsRepo eq.ExerciseQuestionsRepo
}

func New(authRepo eq.ExerciseQuestionsRepo) ExerciseQuestionsUsecase {
	return &exerciseQuestionsUsecase{
		authRepo,
	}
}

package questions

import mc "csw-golang/internal/domain/usecase/exercise/questions"

type ExerciseQuestionsHandler struct {
	exerciseQuestionsUsecase mc.ExerciseQuestionsUsecase
}

func New(exerciseQuestionsUsecase mc.ExerciseQuestionsUsecase) *ExerciseQuestionsHandler {
	return &ExerciseQuestionsHandler{
		exerciseQuestionsUsecase,
	}
}

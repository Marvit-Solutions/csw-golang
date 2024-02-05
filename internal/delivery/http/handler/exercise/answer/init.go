package answer

import ms "csw-golang/internal/domain/usecase/exercise/answer"

type ExerciseAnswerHandler struct {
	exerciseAnswerUsecase ms.ExerciseAnswerUsecase
}

func New(exerciseAnswerUsecase ms.ExerciseAnswerUsecase) *ExerciseAnswerHandler {
	return &ExerciseAnswerHandler{
		exerciseAnswerUsecase,
	}
}

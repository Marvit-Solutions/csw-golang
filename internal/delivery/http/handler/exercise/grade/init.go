package grade

import ms "csw-golang/internal/domain/usecase/exercise/grade"

type ExerciseGradeHandler struct {
	exerciseAnswerUsecase ms.ExerciseGradeUsecase
}

func New(exerciseGradeUsecase ms.ExerciseGradeUsecase) *ExerciseGradeHandler {
	return &ExerciseGradeHandler{
		exerciseGradeUsecase,
	}
}

package submission

import (
	ms "csw-golang/internal/domain/usecase/exercise/submission"
)

type SubmissionHandler struct {
	submissionUsecase ms.ExerciseSubmissionUsecase
}

func New(submissionUsecase ms.ExerciseSubmissionUsecase) *SubmissionHandler {
	return &SubmissionHandler{
		submissionUsecase,
	}
}

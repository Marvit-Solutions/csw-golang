package submission

import (
	"csw-golang/internal/domain/entity/dto"
	es "csw-golang/internal/domain/repository/exercise/submission"
)

type ExerciseSubmissionUsecase interface {
	AddSubmissionUsecase(testTypeID, userID string, answer dto.UserSubmittedAnswerExercisesRequest) (dto.UserTestSubmissionExercisesResponse, error)
	GetSubmissionUsecase(submissionId string) (dto.UserTestSubmissionExercisesResponse, error)
	GetAllSubmissionsUsecase() ([]dto.UserTestSubmissionExercisesResponse, error)
}

type exerciseSubmissionUsecase struct {
	exerciseSubmissionRepo es.ExerciseSubmissionRepo
}

func New(exerciseSubmissionRepo es.ExerciseSubmissionRepo) ExerciseSubmissionUsecase {
	return &exerciseSubmissionUsecase{
		exerciseSubmissionRepo,
	}
}

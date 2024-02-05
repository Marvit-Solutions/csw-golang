package submission

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"github.com/google/uuid"
	"time"
)

func (e exerciseSubmissionUsecase) AddSubmissionUsecase(submission dto.UserTestSubmissionExercisesRequest) (dto.UserTestSubmissionExercisesResponse, error) {

	exercises := datastruct.UserTestSubmissionExercises{
		ID:                 uuid.New().String(),
		UserID:             submission.UserID,
		TestTypeExerciseID: submission.TestTypeExerciseID,
		SubmissionTIme:     time.Now(),
	}

	addSubmission, err := e.exerciseSubmissionRepo.AddSubmission(exercises)

	if err != nil {
		return dto.UserTestSubmissionExercisesResponse{}, err
	}

	response := dto.UserTestSubmissionExercisesResponse{
		ID:                 addSubmission.ID,
		UserID:             addSubmission.UserID,
		TestTypeExerciseID: addSubmission.TestTypeExerciseID,
		SubmissionTime:     addSubmission.SubmissionTIme,
	}

	return response, nil
}

func (e exerciseSubmissionUsecase) GetSubmissionUsecase(submissionId string) (dto.UserTestSubmissionExercisesResponse, error) {
	submission, err := e.exerciseSubmissionRepo.GetSubmission(submissionId)
	if err != nil {
		return dto.UserTestSubmissionExercisesResponse{}, err
	}
	response := dto.UserTestSubmissionExercisesResponse{
		ID:                 submission.ID,
		UserID:             submission.UserID,
		TestTypeExerciseID: submission.TestTypeExerciseID,
		SubmissionTime:     submission.SubmissionTIme,
	}
	return response, nil
}

func (e exerciseSubmissionUsecase) GetAllSubmissionsUsecase() ([]dto.UserTestSubmissionExercisesResponse, error) {

	submissions, err := e.exerciseSubmissionRepo.GetAllSubmissions()
	if err != nil {
		return []dto.UserTestSubmissionExercisesResponse{}, err
	}
	var resArr []dto.UserTestSubmissionExercisesResponse
	for _, sub := range submissions {
		response := dto.UserTestSubmissionExercisesResponse{
			ID:                 sub.ID,
			UserID:             sub.UserID,
			TestTypeExerciseID: sub.TestTypeExerciseID,
			SubmissionTime:     sub.SubmissionTIme,
		}
		resArr = append(resArr, response)
	}
	return resArr, nil

}

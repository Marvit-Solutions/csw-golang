package grade

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"github.com/google/uuid"
	"time"
)

func (e exerciseGradeUsecase) AddGradeUsecase(answer dto.GradeExercisesRequest) (dto.GradeExercisesResponse, error) {

	exercises := datastruct.GradeExercises{
		ID:                           uuid.New().String(),
		UserTestSubmissionExerciseID: answer.UserTestSubmissionExerciseID,
		UserID:                       answer.UserID,
		TestTypeExerciseID:           answer.TestTypeExerciseID,
		GradingTime:                  time.Now(),
	}
	grade, err := e.exerciseGradeRepo.AddGrade(exercises)
	if err != nil {
		return dto.GradeExercisesResponse{}, err
	}

	response := dto.GradeExercisesResponse{
		ID:                           grade.ID,
		UserTestSubmissionExerciseID: grade.UserTestSubmissionExerciseID,
		UserID:                       grade.UserID,
		TestTypeExerciseID:           grade.TestTypeExerciseID,
		Score:                        grade.Score,
		GradingTime:                  grade.GradingTime,
	}

	return response, err
}

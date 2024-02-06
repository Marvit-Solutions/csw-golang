package grade

import (
	"csw-golang/internal/domain/entity/dto"
	es "csw-golang/internal/domain/repository/exercise/grade"
)

type ExerciseGradeUsecase interface {
	AddGradeUsecase(grade dto.GradeExercisesRequest) (dto.GradeExercisesResponse, error)
	GetGradeUsecase(submissionID string, userID string) ([]dto.GradeExercisesResponse, error)
	GetGradeByIdUsecase(gradeId string) (dto.GradeExercisesResponse, error)
}

type exerciseGradeUsecase struct {
	exerciseGradeRepo es.ExerciseGradeRepo
}

func (e exerciseGradeUsecase) GetGradeUsecase(submissionID string, userID string) ([]dto.GradeExercisesResponse, error) {
	grade, err := e.exerciseGradeRepo.GetGrade(userID, submissionID)
	if err != nil {
		return []dto.GradeExercisesResponse{}, err
	}

	var responseArr []dto.GradeExercisesResponse
	for _, gg := range grade {
		response := dto.GradeExercisesResponse{
			ID:                           gg.ID,
			UserTestSubmissionExerciseID: gg.UserTestSubmissionExerciseID,
			UserID:                       gg.UserID,
			TestTypeExerciseID:           gg.TestTypeExerciseID,
			Score:                        gg.Score,
			GradingTime:                  gg.GradingTime,
		}
		responseArr = append(responseArr, response)
	}
	return responseArr, nil

}

func (e exerciseGradeUsecase) GetGradeByIdUsecase(gradeId string) (dto.GradeExercisesResponse, error) {
	gg, err := e.exerciseGradeRepo.GetGradeByID(gradeId)
	if err != nil {
		return dto.GradeExercisesResponse{}, err
	}
	response := dto.GradeExercisesResponse{
		ID:                           gg.ID,
		UserTestSubmissionExerciseID: gg.UserTestSubmissionExerciseID,
		UserID:                       gg.UserID,
		TestTypeExerciseID:           gg.TestTypeExerciseID,
		Score:                        gg.Score,
		GradingTime:                  gg.GradingTime,
	}
	return response, nil
}

func New(exerciseGradeRepo es.ExerciseGradeRepo) ExerciseGradeUsecase {
	return &exerciseGradeUsecase{
		exerciseGradeRepo,
	}
}

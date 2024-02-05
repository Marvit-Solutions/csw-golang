package answer

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"github.com/google/uuid"
)

func (e exerciseAnswerUsecase) AddAnswerUsecase(answer dto.UserSubmittedAnswerExercisesRequest) (dto.UserSubmittedAnswerExercisesResponse, error) {

	var dataArr []datastruct.UserSubmittedAnswerExercises

	for _, aa := range answer.PairOfUserAnswer {
		exercises := datastruct.UserSubmittedAnswerExercises{
			ID:                           uuid.New().String(),
			UserTestSubmissionExerciseID: answer.UserTestSubmissionExerciseID,
			QuestionExerciseID:           aa.QuestionExerciseID,
			ChoiceExerciseID:             aa.ChoiceExerciseID,
		}
		dataArr = append(dataArr, exercises)
	}
	_, err := e.exerciseAnswerRepo.AddAnswer(dataArr)

	if err != nil {
		return dto.UserSubmittedAnswerExercisesResponse{}, err
	}

	return dto.UserSubmittedAnswerExercisesResponse{}, nil
}

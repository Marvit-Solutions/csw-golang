package questions

import (
	"csw-golang/internal/domain/entity/dto"
)

func (e exerciseQuestionsUsecase) AddExerciseQuestion(exerciseQuestion dto.QuestionExercisesRequest) (dto.QuestionExercisesResponse, error) {

	question, err := e.exerciseQuestionsRepo.AddExerciseQuestion(exerciseQuestion)
	if err != nil {
		return dto.QuestionExercisesResponse{}, nil
	}

	var choiceExercisesArray []dto.ChoiceExercisesRequest
	for _, request := range question.ChoiceExercises {
		choiceExercises := dto.ChoiceExercisesRequest{
			Type:      request.Type,
			Content:   request.Content,
			IsCorrect: request.IsCorrect,
			Weight:    request.Weight,
		}

		choiceExercisesArray = append(choiceExercisesArray, choiceExercises)
	}
	response := dto.QuestionExercisesResponse{
		TestTypeExerciseID: question.TestTypeExerciseID,
		Content:            question.Content,
		Weight:             question.Weight,
		//Tags:               question.Tags,
		ChoiceExercises: choiceExercisesArray,
	}

	return response, nil
}

func (e exerciseQuestionsUsecase) GetExerciseQuestions(exerciseQuestionsID string) (dto.QuestionExercisesResponse, error) {

	question, err := e.exerciseQuestionsRepo.GetExerciseQuestions(exerciseQuestionsID)
	if err != nil {
		return dto.QuestionExercisesResponse{}, nil
	}
	var choiceExercisesArray []dto.ChoiceExercisesRequest
	for _, request := range question.ChoiceExercises {
		choiceExercises := dto.ChoiceExercisesRequest{
			Type:      request.Type,
			Content:   request.Content,
			IsCorrect: request.IsCorrect,
			Weight:    request.Weight,
		}

		choiceExercisesArray = append(choiceExercisesArray, choiceExercises)
	}
	response := dto.QuestionExercisesResponse{
		TestTypeExerciseID: question.TestTypeExerciseID,
		Content:            question.Content,
		Weight:             question.Weight,
		//Tags:               question.Tags,
		ChoiceExercises: choiceExercisesArray,
	}

	return response, nil
}

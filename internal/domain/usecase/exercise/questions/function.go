package questions

import (
	"csw-golang/internal/domain/entity/dto"
)

func (e exerciseQuestionsUsecase) AddExerciseQuestion(exerciseQuestion dto.QuestionExercisesRequest) (dto.QuestionExercisesResponse, error) {

	question, err := e.exerciseQuestionsRepo.AddExerciseQuestion(exerciseQuestion)
	if err != nil {
		return dto.QuestionExercisesResponse{}, err
	}

	var choiceExercisesArray []dto.ChoiceExercisesRequest
	for _, request := range question.ChoiceExercises {
		choiceExercises := dto.ChoiceExercisesRequest{
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

func (e exerciseQuestionsUsecase) AddBatchExerciseQuestion(exerciseQuestion []dto.QuestionExercisesRequest) ([]dto.QuestionExercisesResponse, error) {

	question, err := e.exerciseQuestionsRepo.AddBatchExerciseQuestion(exerciseQuestion)
	if err != nil {
		return []dto.QuestionExercisesResponse{}, err
	}
	var dtoExerciseArray []dto.QuestionExercisesResponse
	for _, question := range question {
		var choiceExercisesArray []dto.ChoiceExercisesRequest
		for _, request := range question.ChoiceExercises {
			choiceExercises := dto.ChoiceExercisesRequest{
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
		dtoExerciseArray = append(dtoExerciseArray, response)
	}
	return dtoExerciseArray, nil
}

func (e exerciseQuestionsUsecase) GetExerciseQuestions(exerciseQuestionsID string) (dto.QuestionExercisesResponse, error) {

	question, err := e.exerciseQuestionsRepo.GetExerciseQuestions(exerciseQuestionsID)
	if err != nil {
		return dto.QuestionExercisesResponse{}, err
	}
	var choiceExercisesArray []dto.ChoiceExercisesRequest
	for _, request := range question.ChoiceExercises {
		choiceExercises := dto.ChoiceExercisesRequest{
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

func (e exerciseQuestionsUsecase) GetAllExerciseQuestions() ([]dto.QuestionExercisesResponse, error) {
	question, err := e.exerciseQuestionsRepo.GetALlExerciseQuestions()
	if err != nil {
		return []dto.QuestionExercisesResponse{}, err
	}
	var dtoExerciseArray []dto.QuestionExercisesResponse
	for _, question := range question {
		var choiceExercisesArray []dto.ChoiceExercisesRequest
		for _, request := range question.ChoiceExercises {
			choiceExercises := dto.ChoiceExercisesRequest{
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
		dtoExerciseArray = append(dtoExerciseArray, response)
	}
	return dtoExerciseArray, nil
}

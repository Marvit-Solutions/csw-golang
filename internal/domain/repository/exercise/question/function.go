package question

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"github.com/google/uuid"
)

func (e exerciseQuestionsRepo) AddExerciseQuestion(exerciseQuestion dto.QuestionExercisesRequest) (datastruct.QuestionExercises, error) {
	var choiceExercisesArray []datastruct.ChoiceExercises
	for _, request := range exerciseQuestion.ChoiceExercises {
		choiceExercises := datastruct.ChoiceExercises{
			ID:                           uuid.NewString(),
			Type:                         request.Type,
			Content:                      request.Content,
			IsCorrect:                    request.IsCorrect,
			Weight:                       request.Weight,
			UserSubmittedAnswerExercises: nil, // Initialize as needed
		}

		choiceExercisesArray = append(choiceExercisesArray, choiceExercises)
	}

	exercises := datastruct.QuestionExercises{
		ID:              uuid.New().String(),
		Content:         exerciseQuestion.Content,
		Weight:          exerciseQuestion.Weight,
		ChoiceExercises: choiceExercisesArray,
	}
	// Create record in the database
	result := e.db.Create(&exercises)
	if result.Error != nil {
		return datastruct.QuestionExercises{}, result.Error
	}
	return exercises, nil
}

func (e exerciseQuestionsRepo) AddBatchExerciseQuestion(exerciseQuestion []dto.QuestionExercisesRequest) ([]datastruct.QuestionExercises, error) {

	var questionsArr []datastruct.QuestionExercises

	for _, request := range exerciseQuestion {
		var choiceExercisesArray []datastruct.ChoiceExercises
		for _, request1 := range request.ChoiceExercises {
			choiceExercises := datastruct.ChoiceExercises{
				ID:                           uuid.NewString(),
				Type:                         request1.Type,
				Content:                      request1.Content,
				IsCorrect:                    request1.IsCorrect,
				Weight:                       request1.Weight,
				UserSubmittedAnswerExercises: nil, // Initialize as needed
			}

			choiceExercisesArray = append(choiceExercisesArray, choiceExercises)
		}

		exercises := datastruct.QuestionExercises{
			ID:              uuid.New().String(),
			Content:         request.Content,
			Weight:          request.Weight,
			ChoiceExercises: choiceExercisesArray,
		}
		questionsArr = append(questionsArr, exercises)
	}

	result := e.db.CreateInBatches(questionsArr, 100)
	if result.Error != nil {
		return []datastruct.QuestionExercises{}, result.Error
	}

	return questionsArr, nil
}

func (e exerciseQuestionsRepo) GetExerciseQuestions(exerciseQuestionsID string) (datastruct.QuestionExercises, error) {
	var model datastruct.QuestionExercises

	// Find record in the database based on exerciseQuestionsID
	result := e.db.Preload("ChoiceExercises").Where("id = ?", exerciseQuestionsID).First(&model)
	if result.Error != nil {
		return datastruct.QuestionExercises{}, result.Error
	}
	return model, nil

}

func (e exerciseQuestionsRepo) GetALlExerciseQuestions() ([]datastruct.QuestionExercises, error) {
	var model []datastruct.QuestionExercises

	// Find record in the database based on exerciseQuestionsID
	result := e.db.Preload("ChoiceExercises").Find(&model)
	if result.Error != nil {
		return []datastruct.QuestionExercises{}, result.Error
	}
	return model, nil

}

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

func (e exerciseQuestionsRepo) GetExerciseQuestions(exerciseQuestionsID string) (datastruct.QuestionExercises, error) {

	var model datastruct.QuestionExercises

	// Find record in the database based on exerciseQuestionsID
	result := e.db.Preload("ChoiceExercises").Where("id = ?", exerciseQuestionsID).First(&model)
	if result.Error != nil {
		return datastruct.QuestionExercises{}, result.Error
	}
	return model, nil
}

package seed

import (
	ds "csw-golang/internal/domain/entity/datastruct"
	"gorm.io/gorm"
	"time"
)

func CreateQuestions() []*ds.QuestionExercises {
	QuestionsExercise := []*ds.QuestionExercises{
		{
			ID:        "d857d9b1-7c37-43d6-9b62-2679ee60778b",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
			Content:   "Calculate the limit as x approaches 0 for (sin(x)/x) ?",
			Weight:    15,
			ChoiceExercises: []ds.ChoiceExercises{{
				ID:                           "e8eb90c3-00fd-45e1-9f7e-bcd45f64d2f6",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Type:                         "A",
				Content:                      "0",
				IsCorrect:                    true,
				Weight:                       15,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "296f1912-c455-442c-9904-43917d258ebf",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Type:                         "B",
				Content:                      "1",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "0f33c05a-f7c7-42bf-b9e5-a61bf6b8c3cd",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Type:                         "C",
				Content:                      "Undefined",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "3c4cda6d-bb4f-4f97-963b-1e397a879200",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Type:                         "D",
				Content:                      "Infinity",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}, {
				ID:                           "3c4cda6d-bb4f-4f97-963b-1e397a879200",
				CreatedAt:                    time.Time{},
				UpdatedAt:                    time.Time{},
				DeletedAt:                    gorm.DeletedAt{},
				Type:                         "E",
				Content:                      "phi",
				IsCorrect:                    false,
				Weight:                       0,
				UserSubmittedAnswerExercises: nil,
			}},
			UserSubmittedAnswerExercises: nil,
		},
	}

	return QuestionsExercise
}

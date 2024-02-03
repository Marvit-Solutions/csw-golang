package question

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"gorm.io/gorm"
)

type ExerciseQuestionsRepo interface {
	AddExerciseQuestion(exerciseQuestion dto.QuestionExercisesRequest) (datastruct.QuestionExercises, error)
	AddBatchExerciseQuestion(exerciseQuestion []dto.QuestionExercisesRequest) ([]datastruct.QuestionExercises, error)
	GetExerciseQuestions(exerciseQuestionsID string) (datastruct.QuestionExercises, error)
	GetALlExerciseQuestions() ([]datastruct.QuestionExercises, error)
}

type exerciseQuestionsRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ExerciseQuestionsRepo {
	return &exerciseQuestionsRepo{
		db,
	}
}

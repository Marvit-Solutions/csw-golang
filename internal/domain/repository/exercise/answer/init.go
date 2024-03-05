package answer

import (
	"csw-golang/internal/domain/entity/datastruct"
	"gorm.io/gorm"
)

type ExerciseAnswerRepo interface {
	AddAnswer(answer []datastruct.UserSubmittedAnswerExercises) ([]datastruct.UserSubmittedAnswerExercises, error)
}

type exercisesAnswerRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ExerciseAnswerRepo {
	return &exercisesAnswerRepo{
		db,
	}
}

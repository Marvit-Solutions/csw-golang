package submission

import (
	"csw-golang/internal/domain/entity/datastruct"
	"gorm.io/gorm"
)

type ExerciseSubmissionRepo interface {
	AddSubmission(submission datastruct.UserTestSubmissionExercises) (datastruct.UserTestSubmissionExercises, error)
	GetSubmission(submissionId string) (datastruct.UserTestSubmissionExercises, error)
	GetAllSubmissions() ([]datastruct.UserTestSubmissionExercises, error)
}

type exercisesSubmissionRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ExerciseSubmissionRepo {
	return &exercisesSubmissionRepo{
		db,
	}
}

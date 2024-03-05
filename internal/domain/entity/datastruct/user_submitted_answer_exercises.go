package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserSubmittedAnswerExercises struct {
	ID                           string         `gorm:"type:text;primaryKey"`
	CreatedAt                    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                    gorm.DeletedAt `gorm:"index"`
	UserTestSubmissionExerciseID string         `json:"UserTestSubmissionExerciseID" form:"UserTestSubmissionExerciseID"`
	QuestionExerciseID           string         `json:"QuestionExerciseID" form:"QuestionExerciseID"`
	ChoiceExerciseID             string         `json:"ChoiceExerciseID" form:"ChoiceExerciseID"`
}

type UserSubmittedAnswerExercises2 struct {
	ID                           string            `gorm:"type:text;primaryKey"`
	CreatedAt                    time.Time         `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                    time.Time         `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                    gorm.DeletedAt    `gorm:"index"`
	UserTestSubmissionExerciseID string            `json:"UserTestSubmissionExerciseID" form:"UserTestSubmissionExerciseID"`
	QuestionExercise             QuestionExercises `json:"QuestionExerciseID" form:"QuestionExerciseID"`
	ChoiceExercise               ChoiceExercises   `json:"ChoiceExerciseID" form:"ChoiceExerciseID"`
}

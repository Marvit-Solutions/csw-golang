package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserTestSubmissionExercises struct {
	ID                           string                         `gorm:"type:text;primaryKey"`
	CreatedAt                    time.Time                      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                    time.Time                      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                    gorm.DeletedAt                 `gorm:"index"`
	UserID                       string                         `json:"UserID" form:"UserID"`
	TestTypeExerciseID           string                         `json:"TestTypeExerciseID" form:"TestTypeExerciseID"`
	SubmissionTIme               time.Time                      `json:"SubmissionTIme" form:"SubmissionTIme"`
	GradeExercises               GradeExercises                 `gorm:"foreignKey:UserTestSubmissionExerciseID"`
	UserSubmittedAnswerExercises []UserSubmittedAnswerExercises `gorm:"foreignKey:UserTestSubmissionExerciseID"`
}

package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type GradeExercises struct {
	ID                           string         `gorm:"type:text;primaryKey"`
	CreatedAt                    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                    time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                    gorm.DeletedAt `gorm:"index"`
	UserTestSubmissionExerciseID string         `json:"UserTestSubmissionExerciseID" form:"UserTestSubmissionExerciseID"`
	UserID                       string         `json:"UserID" form:"UserID"`
	TestTypeExerciseID           string         `json:"TestTypeExerciseID" form:"TestTypeExerciseID"`
	Score                        int            `json:"Score" form:"Score"`
	GradingTime                  time.Time      `json:"GradingTime" form:"GradingTime"`
}

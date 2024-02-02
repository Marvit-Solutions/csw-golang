package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type GradeQuizzes struct {
	ID                       string         `gorm:"type:text;primaryKey"`
	CreatedAt                time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                gorm.DeletedAt `gorm:"index"`
	UserTestSubmissionQuizID string         `json:"UserTestSubmissionQuizID" form:"UserTestSubmissionQuizID"`
	UserID                   string         `json:"UserID" form:"UserID"`
	TestTypeQuizID           string         `json:"TestTypeQuizID" form:"TestTypeQuizID"`
	Score                    int            `json:"Score" form:"Score"`
	GradingTime              time.Time      `json:"GradingTime" form:"GradingTime"`
}

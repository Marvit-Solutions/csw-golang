package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserSubmittedAnswerQuizzes struct {
	ID                       string         `gorm:"type:text;primaryKey"`
	CreatedAt                time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                gorm.DeletedAt `gorm:"index"`
	UserTestSubmissionQuizID string         `json:"UserTestSubmissionQuizID" form:"UserTestSubmissionQuizID"`
	QuestionQuizID           string         `json:"QuestionQuizID" form:"QuestionQuizID"`
	ChoiceQuizID             string         `json:"ChoiceQuizID" form:"ChoiceQuizID"`
}

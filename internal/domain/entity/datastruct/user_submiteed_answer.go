package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserSubmittedAnswers struct {
	ID                   string         `gorm:"type:text;primaryKey"`
	CreatedAt            time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	UserTestSubmissionID string         `json:"UserTestSubmissionID" form:"UserTestSubmissionID"`
	QuestionID           string         `json:"QuestionID" form:"QuestionID"`
	ChoiceID             string         `json:"ChoiceID" form:"ChoiceID"`
}

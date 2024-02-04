package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type ChoiceQuizzes struct {
	ID                         string                       `gorm:"type:text;primaryKey"`
	CreatedAt                  time.Time                    `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                  time.Time                    `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                  gorm.DeletedAt               `gorm:"index"`
	QuestionQuizID             string                       `json:"QuestionQuizID" form:"QuestionQuizID"`
	Content                    string                       `json:"Content" form:"Content"`
	IsCorrect                  bool                         `json:"IsCorrect" form:"IsCorrect"`
	Weight                     int                          `json:"Weight" form:"Weight"`
	UserSubmittedAnswerQuizzes []UserSubmittedAnswerQuizzes `gorm:"foreignKey:ChoiceQuizID"`
}

package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type QuestionQuizzes struct {
	ID                         string                       `gorm:"type:text;primaryKey"`
	CreatedAt                  time.Time                    `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                  time.Time                    `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                  gorm.DeletedAt               `gorm:"index"`
	TestTypeQuizID             string                       `json:"TestTypeQuizID" form:"TestTypeQuizID"`
	Content                    string                       `json:"Content" form:"Content"`
	Weight                     int                          `json:"Weight" form:"Weight"`
	ChoiceQuizzes              []ChoiceQuizzes              `gorm:"foreignKey:QuestionQuizID"`
	UserSubmittedAnswerQuizzes []UserSubmittedAnswerQuizzes `gorm:"foreignKey:QuestionQuizID"`
}

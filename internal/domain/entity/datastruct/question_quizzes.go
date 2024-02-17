package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type QuestionQuizzes struct {
	ID             string         `gorm:"type:text;primaryKey"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	TestTypeQuizID string         `json:"TestTypeQuizID" form:"TestTypeQuizID"`
	Image          string         `json:"Image" form:"Image" gorm:"default: 'assets/img/questions/question.png'"`
	Content        string         `json:"Content" form:"Content"`
	Weight         int            `json:"Weight" form:"Weight"`
	// Status                     string                       `json:"Status" form:"Status"  gorm:"default: 'Belum Dijawab'"`
	ChoiceQuizzes              []ChoiceQuizzes              `gorm:"foreignKey:QuestionQuizID"`
	UserSubmittedAnswerQuizzes []UserSubmittedAnswerQuizzes `gorm:"foreignKey:QuestionQuizID"`
}

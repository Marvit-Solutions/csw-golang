package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Choices struct {
	ID                  string                 `gorm:"type:text;primaryKey"`
	CreatedAt           time.Time              `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time              `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt         `gorm:"index"`
	QuestionID          string                 `json:"QuestionID" form:"QuestionID"`
	Type                string                 `json:"Type" form:"Type"`
	Content             string                 `json:"Content" form:"Content"`
	IsCorrect           bool                   `json:"IsCorrect" form:"IsCorrect"`
	Weight              int                    `json:"Weight" form:"Weight"`
	UserSubmittedAnswer []UserSubmittedAnswers `gorm:"foreignKey:ChoiceID"`
}

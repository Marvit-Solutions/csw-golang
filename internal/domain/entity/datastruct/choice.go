package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Choice struct {
	ID                  string                `gorm:"type:text;primaryKey"`
	CreatedAt           time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt        `gorm:"index"`
	QuestionID          string                `json:"QuestionID" form:"QuestionID"`
	Type                string                `json:"Type" form:"Type"`
	Content             string                `json:"Content" form:"Content"`
	Is_correct          bool                  `json:"Is_correct" form:"Is_correct"`
	Weight              int                   `json:"Weight" form:"Weight"`
	UserSubmittedAnswer []UserSubmittedAnswer `gorm:"foreignKey:ChoiceID"`
}

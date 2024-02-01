package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID                  string                `gorm:"type:text;primaryKey"`
	CreatedAt           time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt        `gorm:"index"`
	TestTypeID          string                `json:"TestTypeID" form:"TestTypeID"`
	Content             string                `json:"Content" form:"Content"`
	Weight              int                   `json:"Weight" form:"Weight"`
	Choice              []Choice              `gorm:"foreignKey:QuestionID"`
	UserSubmittedAnswer []UserSubmittedAnswer `gorm:"foreignKey:QuestionID"`
}

package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserTestSubmissions struct {
	ID                   string                 `gorm:"type:text;primaryKey"`
	CreatedAt            time.Time              `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time              `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt            gorm.DeletedAt         `gorm:"index"`
	UserID               string                 `json:"UserID" form:"UserID"`
	TestTypeID           string                 `json:"TestTypeID" form:"TestTypeID"`
	SubmissionTIme       time.Time              `json:"SubmissionTIme" form:"SubmissionTIme"`
	Grades               Grades                 `gorm:"foreignKey:UserTestSubmissionID"`
	UserSubmittedAnswers []UserSubmittedAnswers `gorm:"foreignKey:UserTestSubmissionID"`
}

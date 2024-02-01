package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserTestSubmission struct {
	ID                  string                `gorm:"type:text;primaryKey"`
	CreatedAt           time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt        `gorm:"index"`
	UserID              string                `json:"UserID" form:"UserID"`
	TestTypeID          string                `json:"TestTypeID" form:"TestTypeID"`
	SubmissionTIme      time.Time             `json:"SubmissionTIme" form:"SubmissionTIme"`
	Grade               Grade                 `gorm:"foreignKey:UserTestSubmissionID"`
	UserSubmittedAnswer []UserSubmittedAnswer `gorm:"foreignKey:UserTestSubmissionID"`
}

package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Grade struct {
	ID                   string         `gorm:"type:text;primaryKey"`
	CreatedAt            time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	UserTestSubmissionID string         `json:"UserTestSubmissionID" form:"UserTestSubmissionID"`
	UserID               string         `json:"UserID" form:"UserID"`
	TestTypeID           string         `json:"TestTypeID" form:"TestTypeID"`
	Score                int            `json:"Score" form:"Score"`
	GradingTime          time.Time      `json:"GradingTime" form:"GradingTime"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQuizSubmission = "quiz_submissions"

// QuizSubmission mapped from table <quiz_submissions>
type QuizSubmission struct {
	ID           string         `gorm:"column:id;primaryKey" json:"id"`
	UserID       string         `gorm:"column:user_id;not null" json:"user_id"`
	QuizID       string         `gorm:"column:quiz_id;not null" json:"quiz_id"`
	StartedAt    time.Time      `gorm:"column:started_at;not null" json:"started_at"`
	FinishedAt   time.Time      `gorm:"column:finished_at;not null;default:CURRENT_TIMESTAMP" json:"finished_at"`
	TimeRequired string         `gorm:"column:time_required;not null" json:"time_required"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName QuizSubmission's table name
func (*QuizSubmission) TableName() string {
	return TableNameQuizSubmission
}

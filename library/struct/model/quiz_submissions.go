package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQuizSubmission = "quiz_submissions"

// QuizSubmission mapped from table <quiz_submissions>
type QuizSubmission struct {
	ID           int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID         string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	UserID       int          `gorm:"column:user_id;not null" json:"user_id"`
	QuizID       int          `gorm:"column:quiz_id;not null" json:"quiz_id"`
	StartedAt    time.Time      `gorm:"column:started_at;not null" json:"started_at"`
	FinishedAt   time.Time      `gorm:"column:finished_at;not null;default:CURRENT_TIMESTAMP" json:"finished_at"`
	TimeRequired string         `gorm:"column:time_required;not null" json:"time_required"`
	RightAnswer  int          `gorm:"column:right_answer;not null" json:"right_answer"`
	Score        int          `gorm:"column:score;not null" json:"score"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName QuizSubmission's table name
func (*QuizSubmission) TableName() string {
	return TableNameQuizSubmission
}

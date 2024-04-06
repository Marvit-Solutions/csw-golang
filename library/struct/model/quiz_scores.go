package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQuizScore = "quiz_scores"

// QuizScore mapped from table <quiz_scores>
type QuizScore struct {
	ID           int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID         string         `gorm:"column:uuid;not null" json:"uuid"`
	SubmissionID int          `gorm:"column:submission_id;not null" json:"submission_id"`
	Score        int          `gorm:"column:score;not null" json:"score"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName QuizScore's table name
func (*QuizScore) TableName() string {
	return TableNameQuizScore
}

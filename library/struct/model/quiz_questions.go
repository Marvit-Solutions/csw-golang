package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQuizQuestion = "quiz_questions"

// QuizQuestion mapped from table <quiz_questions>
type QuizQuestion struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	QuizID    string         `gorm:"column:quiz_id;not null" json:"quiz_id"`
	Content   string         `gorm:"column:content;not null" json:"content"`
	Image     *string        `gorm:"column:image" json:"image"`
	Point     int          `gorm:"column:point;not null" json:"point"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName QuizQuestion's table name
func (*QuizQuestion) TableName() string {
	return TableNameQuizQuestion
}

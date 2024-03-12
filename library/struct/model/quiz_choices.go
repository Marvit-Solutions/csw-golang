package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQuizChoice = "quiz_choices"

// QuizChoice mapped from table <quiz_choices>
type QuizChoice struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	QuestionID string         `gorm:"column:question_id;not null" json:"question_id"`
	Content    string         `gorm:"column:content;not null" json:"content"`
	Point      int          `gorm:"column:point;not null" json:"point"`
	IsCorrect  bool           `gorm:"column:is_correct;not null" json:"is_correct"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName QuizChoice's table name
func (*QuizChoice) TableName() string {
	return TableNameQuizChoice
}

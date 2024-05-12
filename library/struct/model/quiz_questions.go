package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQuizQuestion = "quiz_questions"

// QuizQuestion mapped from table <quiz_questions>
type QuizQuestion struct {
	ID          int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID        string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	QuizID      int          `gorm:"column:quiz_id;not null" json:"quiz_id"`
	Content     string         `gorm:"column:content;not null" json:"content"`
	Score       int          `gorm:"column:score;not null" json:"score"`
	Explanation string         `gorm:"column:explanation;not null" json:"explanation"`
	CreatedBy   int          `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy   int          `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName QuizQuestion's table name
func (*QuizQuestion) TableName() string {
	return TableNameQuizQuestion
}

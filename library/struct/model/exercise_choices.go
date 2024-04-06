package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExerciseChoice = "exercise_choices"

// ExerciseChoice mapped from table <exercise_choices>
type ExerciseChoice struct {
	ID         int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID       string         `gorm:"column:uuid;not null" json:"uuid"`
	QuestionID int          `gorm:"column:question_id;not null" json:"question_id"`
	Content    string         `gorm:"column:content;not null" json:"content"`
	Point      string         `gorm:"column:point;not null" json:"point"`
	IsCorrect  bool           `gorm:"column:is_correct;not null" json:"is_correct"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ExerciseChoice's table name
func (*ExerciseChoice) TableName() string {
	return TableNameExerciseChoice
}

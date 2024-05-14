package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExerciseChoice = "exercise_choices"

// ExerciseChoice mapped from table <exercise_choices>
type ExerciseChoice struct {
	ID         int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID       string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	QuestionID int          `gorm:"column:question_id;not null" json:"question_id"`
	Content    string         `gorm:"column:content;not null" json:"content"`
	Score      int          `gorm:"column:score;not null" json:"score"`
	IsCorrect  bool           `gorm:"column:is_correct;not null" json:"is_correct"`
	CreatedBy  int          `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy  int          `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ExerciseChoice's table name
func (*ExerciseChoice) TableName() string {
	return TableNameExerciseChoice
}

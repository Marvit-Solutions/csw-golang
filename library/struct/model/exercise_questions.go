package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExerciseQuestion = "exercise_questions"

// ExerciseQuestion mapped from table <exercise_questions>
type ExerciseQuestion struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	ExerciseID string         `gorm:"column:exercise_id;not null" json:"exercise_id"`
	Content    string         `gorm:"column:content;not null" json:"content"`
	Image      *string        `gorm:"column:image" json:"image"`
	Point      int          `gorm:"column:point;not null" json:"point"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ExerciseQuestion's table name
func (*ExerciseQuestion) TableName() string {
	return TableNameExerciseQuestion
}

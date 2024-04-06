package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExerciseScore = "exercise_scores"

// ExerciseScore mapped from table <exercise_scores>
type ExerciseScore struct {
	ID           int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID         string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	SubmissionID int          `gorm:"column:submission_id;not null" json:"submission_id"`
	Score        int          `gorm:"column:score;not null" json:"score"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ExerciseScore's table name
func (*ExerciseScore) TableName() string {
	return TableNameExerciseScore
}

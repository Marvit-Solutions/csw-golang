package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExerciseSubmission = "exercise_submissions"

// ExerciseSubmission mapped from table <exercise_submissions>
type ExerciseSubmission struct {
	ID           int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID         string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	UserID       int          `gorm:"column:user_id;not null" json:"user_id"`
	ExerciseID   int          `gorm:"column:exercise_id;not null" json:"exercise_id"`
	StartedAt    time.Time      `gorm:"column:started_at;not null" json:"started_at"`
	FinishedAt   time.Time      `gorm:"column:finished_at;not null" json:"finished_at"`
	TimeRequired string         `gorm:"column:time_required;not null" json:"time_required"`
	Score        int          `gorm:"column:score;not null" json:"score"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ExerciseSubmission's table name
func (*ExerciseSubmission) TableName() string {
	return TableNameExerciseSubmission
}

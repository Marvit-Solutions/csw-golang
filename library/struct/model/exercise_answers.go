package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExerciseAnswer = "exercise_answers"

// ExerciseAnswer mapped from table <exercise_answers>
type ExerciseAnswer struct {
	ID           int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID         string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	SubmissionID int          `gorm:"column:submission_id;not null" json:"submission_id"`
	ChoiceID     *int         `gorm:"column:choice_id" json:"choice_id"`
	IsMarked     bool           `gorm:"column:is_marked;not null" json:"is_marked"`
	CreatedBy    int          `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy    int          `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ExerciseAnswer's table name
func (*ExerciseAnswer) TableName() string {
	return TableNameExerciseAnswer
}

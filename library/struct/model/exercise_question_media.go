package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExerciseQuestionMedia = "exercise_question_media"

// ExerciseQuestionMedia mapped from table <exercise_question_media>
type ExerciseQuestionMedia struct {
	ID                 int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID               string         `gorm:"column:uuid;not null" json:"uuid"`
	ExerciseQuestionID int            `gorm:"column:exercise_question_id;not null" json:"exercise_question_id"`
	MediaID            int            `gorm:"column:media_id;not null" json:"media_id"`
	Index              int            `gorm:"column:index;not null" json:"index"`
	CreatedAt          time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ExerciseQuestionMedia's table name
func (*ExerciseQuestionMedia) TableName() string {
	return TableNameExerciseQuestionMedia
}

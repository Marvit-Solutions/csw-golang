package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQuizQuestionMedia = "quiz_question_media"

// QuizQuestionMedia mapped from table <quiz_question_media>
type QuizQuestionMedia struct {
	ID             int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID           string         `gorm:"column:uuid;not null" json:"uuid"`
	QuizQuestionID int            `gorm:"column:quiz_question_id;not null" json:"quiz_question_id"`
	MediaID        int            `gorm:"column:media_id;not null" json:"media_id"`
	Index          int            `gorm:"column:index;not null" json:"index"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName QuizQuestionMedia's table name
func (*QuizQuestionMedia) TableName() string {
	return TableNameQuizQuestionMedia
}

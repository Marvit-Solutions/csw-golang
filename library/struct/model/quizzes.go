package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQuiz = "quizzes"

// Quiz mapped from table <quizzes>
type Quiz struct {
	ID           string         `gorm:"column:id;primaryKey" json:"id"`
	SubSubjectID string         `gorm:"column:sub_subject_id;not null" json:"sub_subject_id"`
	TestTypeID   string         `gorm:"column:test_type_id;not null" json:"test_type_id"`
	Open         time.Time      `gorm:"column:open;not null" json:"open"`
	Title        string         `gorm:"column:title;not null" json:"title"`
	Description  string         `gorm:"column:description;not null" json:"description"`
	Time         int          `gorm:"column:time;not null" json:"time"`
	Point        int          `gorm:"column:point;not null" json:"point"`
	Attempt      int          `gorm:"column:attempt;not null" json:"attempt"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName Quiz's table name
func (*Quiz) TableName() string {
	return TableNameQuiz
}

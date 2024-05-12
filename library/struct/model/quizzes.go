package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameQuiz = "quizzes"

// Quiz mapped from table <quizzes>
type Quiz struct {
	ID           int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID         string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	SubSubjectID int          `gorm:"column:sub_subject_id;not null" json:"sub_subject_id"`
	TestTypeID   int          `gorm:"column:test_type_id;not null" json:"test_type_id"`
	Open         time.Time      `gorm:"column:open;not null" json:"open"`
	Title        string         `gorm:"column:title;not null" json:"title"`
	Description  string         `gorm:"column:description;not null" json:"description"`
	Time         int          `gorm:"column:time;not null" json:"time"`
	MaxScore     int          `gorm:"column:max_score;not null" json:"max_score"`
	Attempt      int          `gorm:"column:attempt;not null" json:"attempt"`
	CreatedBy    int          `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy    int          `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Quiz's table name
func (*Quiz) TableName() string {
	return TableNameQuiz
}

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExercise = "exercises"

// Exercise mapped from table <exercises>
type Exercise struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	TestTypeID  string         `gorm:"column:test_type_id;not null" json:"test_type_id"`
	Title       string         `gorm:"column:title;not null" json:"title"`
	Description string         `gorm:"column:description;not null" json:"description"`
	Time        int          `gorm:"column:time;not null" json:"time"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName Exercise's table name
func (*Exercise) TableName() string {
	return TableNameExercise
}

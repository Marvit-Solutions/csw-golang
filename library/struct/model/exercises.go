package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameExercise = "exercises"

// Exercise mapped from table <exercises>
type Exercise struct {
	ID            int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID          string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	TestTypeID    int            `gorm:"column:test_type_id;not null" json:"test_type_id"`
	ModuleID      int            `gorm:"column:module_id;not null" json:"module_id"`
	Title         string         `gorm:"column:title;not null" json:"title"`
	TotalQuestion int            `gorm:"column:total_question;not null" json:"total_question"`
	Description   string         `gorm:"column:description;not null" json:"description"`
	Time          int            `gorm:"column:time;not null" json:"time"`
	Attempt       int            `gorm:"column:attempt;not null" json:"attempt"`
	CreatedBy     int            `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy     int            `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Exercise's table name
func (*Exercise) TableName() string {
	return TableNameExercise
}

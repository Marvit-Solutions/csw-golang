package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSubject = "subjects"

// Subject mapped from table <subjects>
type Subject struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	SubModuleID string         `gorm:"column:sub_module_id;not null" json:"sub_module_id"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName Subject's table name
func (*Subject) TableName() string {
	return TableNameSubject
}

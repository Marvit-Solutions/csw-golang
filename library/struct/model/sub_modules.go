package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSubModule = "sub_modules"

// SubModule mapped from table <sub_modules>
type SubModule struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	ModuleID    string         `gorm:"column:module_id;not null" json:"module_id"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	Description string         `gorm:"column:description;not null" json:"description"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName SubModule's table name
func (*SubModule) TableName() string {
	return TableNameSubModule
}

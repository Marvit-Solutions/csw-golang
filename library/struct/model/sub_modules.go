package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSubModule = "sub_modules"

// SubModule mapped from table <sub_modules>
type SubModule struct {
	ID          int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID        string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	ModuleID    int          `gorm:"column:module_id;not null" json:"module_id"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	Slug        string         `gorm:"column:slug;not null" json:"slug"`
	Description string         `gorm:"column:description;not null" json:"description"`
	CreatedBy   int          `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy   int          `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName SubModule's table name
func (*SubModule) TableName() string {
	return TableNameSubModule
}

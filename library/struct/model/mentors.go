package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMentor = "mentors"

// Mentor mapped from table <mentors>
type Mentor struct {
	ID          int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID        string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	UserID      int          `gorm:"column:user_id;not null" json:"user_id"`
	ModuleID    int          `gorm:"column:module_id;not null" json:"module_id"`
	ShortName   string         `gorm:"column:short_name;not null" json:"short_name"`
	Description string         `gorm:"column:description;not null" json:"description"`
	Motto       string         `gorm:"column:motto;not null" json:"motto"`
	CreatedBy   int          `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy   int          `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Mentor's table name
func (*Mentor) TableName() string {
	return TableNameMentor
}

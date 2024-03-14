package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameRole = "roles"

// Role mapped from table <roles>
type Role struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Role's table name
func (*Role) TableName() string {
	return TableNameRole
}

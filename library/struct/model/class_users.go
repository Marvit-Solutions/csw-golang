package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameClassUser = "class_users"

// ClassUser mapped from table <class_users>
type ClassUser struct {
	ID        int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID      string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	Slug      string         `gorm:"column:slug;not null" json:"slug"`
	CreatedBy int          `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy int          `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ClassUser's table name
func (*ClassUser) TableName() string {
	return TableNameClassUser
}

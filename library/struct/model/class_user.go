package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameClassUser = "class_user"

// ClassUser mapped from table <class_user>
type ClassUser struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName ClassUser's table name
func (*ClassUser) TableName() string {
	return TableNameClassUser
}

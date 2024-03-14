package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUnique = "uniques"

// Unique mapped from table <uniques>
type Unique struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	MentorID  string         `gorm:"column:mentor_id;not null" json:"mentor_id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Unique's table name
func (*Unique) TableName() string {
	return TableNameUnique
}

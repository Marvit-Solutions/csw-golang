package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMentor = "mentors"

// Mentor mapped from table <mentors>
type Mentor struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	UserID      string         `gorm:"column:user_id;not null" json:"user_id"`
	ShortName   string         `gorm:"column:short_name;not null" json:"short_name"`
	Type        string         `gorm:"column:type;not null" json:"type"`
	Description string         `gorm:"column:description;not null" json:"description"`
	Motto       string         `gorm:"column:motto;not null" json:"motto"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Mentor's table name
func (*Mentor) TableName() string {
	return TableNameMentor
}

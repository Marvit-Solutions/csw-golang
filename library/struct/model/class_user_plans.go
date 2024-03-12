package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameClassUserPlan = "class_user_plans"

// ClassUserPlan mapped from table <class_user_plans>
type ClassUserPlan struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	UserID    string         `gorm:"column:user_id;not null" json:"user_id"`
	PlanID    string         `gorm:"column:plan_id;not null" json:"plan_id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName ClassUserPlan's table name
func (*ClassUserPlan) TableName() string {
	return TableNameClassUserPlan
}

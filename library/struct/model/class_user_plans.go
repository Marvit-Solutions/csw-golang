package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameClassUserPlan = "class_user_plans"

// ClassUserPlan mapped from table <class_user_plans>
type ClassUserPlan struct {
	ID        int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID      string         `gorm:"column:uuid;not null" json:"uuid"`
	UserID    int          `gorm:"column:user_id;not null" json:"user_id"`
	PlanID    int          `gorm:"column:plan_id;not null" json:"plan_id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName ClassUserPlan's table name
func (*ClassUserPlan) TableName() string {
	return TableNameClassUserPlan
}

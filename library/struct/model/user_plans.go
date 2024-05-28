package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserPlan = "user_plans"

// UserPlan mapped from table <user_plans>
type UserPlan struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID      string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	UserID    int            `gorm:"column:user_id;not null" json:"user_id"`
	PlanID    int            `gorm:"column:plan_id;not null" json:"plan_id"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName UserPlan's table name
func (*UserPlan) TableName() string {
	return TableNameUserPlan
}

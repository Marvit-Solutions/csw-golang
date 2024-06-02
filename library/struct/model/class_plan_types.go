package model

import (
	"time"
)

const TableNameClassPlanType = "class_plan_types"

// ClassPlanType mapped from table <class_plan_types>
type ClassPlanType struct {
	ID        int      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID      string     `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	PlanID    int      `gorm:"column:plan_id;not null" json:"plan_id"`
	Name      string     `gorm:"column:name;not null" json:"name"`
	Slug      string     `gorm:"column:slug;not null" json:"slug"`
	CreatedAt time.Time  `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_At" json:"deleted_At"`
}

// TableName ClassPlanType's table name
func (*ClassPlanType) TableName() string {
	return TableNameClassPlanType
}

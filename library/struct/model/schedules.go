package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSchedule = "schedules"

// Schedule mapped from table <schedules>
type Schedule struct {
	ID           int            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID         string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	SubSubjectID int            `gorm:"column:sub_subject_id;not null" json:"sub_subject_id"`
	UserPlanID   int            `gorm:"column:user_plan_id;not null" json:"user_plan_id"`
	MeetingDate  time.Time      `gorm:"column:meeting_date;not null" json:"meeting_date"`
	CreatedBy    int            `gorm:"column:created_by;not null" json:"created_by"`
	UpdatedBy    int            `gorm:"column:updated_by;not null" json:"updated_by"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Schedule's table name
func (*Schedule) TableName() string {
	return TableNameSchedule
}

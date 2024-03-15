package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSchedule = "schedules"

// Schedule mapped from table <schedules>
type Schedule struct {
	ID              string         `gorm:"column:id;primaryKey" json:"id"`
	SubSubjectID    string         `gorm:"column:sub_subject_id;not null" json:"sub_subject_id"`
	ClassUserPlanID string         `gorm:"column:class_user_plan_id;not null" json:"class_user_plan_id"`
	MeetingDate     time.Time      `gorm:"column:meeting_date;not null" json:"meeting_date"`
	CreatedAt       time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Schedule's table name
func (*Schedule) TableName() string {
	return TableNameSchedule
}

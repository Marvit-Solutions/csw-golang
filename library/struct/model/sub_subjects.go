package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSubSubject = "sub_subjects"

// SubSubject mapped from table <sub_subjects>
type SubSubject struct {
	ID        int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID      string         `gorm:"column:uuid;not null" json:"uuid"`
	SubjectID int          `gorm:"column:subject_id;not null" json:"subject_id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	Content   string         `gorm:"column:content;not null" json:"content"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName SubSubject's table name
func (*SubSubject) TableName() string {
	return TableNameSubSubject
}

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSubSubjectMedium = "sub_subject_media"

// SubSubjectMedium mapped from table <sub_subject_media>
type SubSubjectMedium struct {
	ID           int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID         string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	SubSubjectID int          `gorm:"column:sub_subject_id;not null" json:"sub_subject_id"`
	MediaID      int          `gorm:"column:media_id;not null" json:"media_id"`
	Index        int          `gorm:"column:index;not null" json:"index"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName SubSubjectMedium's table name
func (*SubSubjectMedium) TableName() string {
	return TableNameSubSubjectMedium
}

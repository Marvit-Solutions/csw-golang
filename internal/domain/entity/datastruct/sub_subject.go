package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubSubject struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	SubjectID string         `json:"SubjectID" form:"SubjectID"`
	Name      string         `json:"Name" form:"Name"`
	Content   string         `json:"Content" form:"Content"`
}

package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Subjects struct {
	ID              string            `gorm:"type:text;primaryKey"`
	CreatedAt       time.Time         `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time         `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt       gorm.DeletedAt    `gorm:"index"`
	SubModuleID     string            `json:"SubModuleID" form:"SubModuleID"`
	Name            string            `json:"Name" form:"Name"`
	SubSubject      []SubSubjects      `gorm:"foreignKey:SubjectID"`
	SubjectTestType []SubjectTestTypes `gorm:"foreignKey:SubjectID"`
}

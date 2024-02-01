package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubjectTestType struct {
	ID                  string                `gorm:"type:text;primaryKey"`
	CreatedAt           time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt        `gorm:"index"`
	SubjectID           string                `json:"SubjectID" form:"SubjectID"`
	TestType            string                `json:"TestType" form:"TestType"`
	Grades              Grades                `gorm:"foreignKey:TestTypeID"`
	Questions           []Questions           `gorm:"foreignKey:TestTypeID"`
	UserTestSubmissions []UserTestSubmissions `gorm:"foreignKey:TestTypeID"`
}

package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubjectTestTypes struct {
	ID                 string                `gorm:"type:text;primaryKey"`
	CreatedAt          time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt          gorm.DeletedAt        `gorm:"index"`
	SubjectID          string                `json:"SubjectID" form:"SubjectID"`
	TestType           string                `json:"TestType" form:"TestType"`
	Grade              Grades                `gorm:"foreignKey:TestTypeID"`
	Question           []Questions           `gorm:"foreignKey:TestTypeID"`
	UserTestSubmission []UserTestSubmissions `gorm:"foreignKey:TestTypeID"`
}

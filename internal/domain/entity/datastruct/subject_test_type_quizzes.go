package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubjectTestTypeQuizzes struct {
	ID                        string                      `gorm:"type:text;primaryKey"`
	CreatedAt                 time.Time                   `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                 time.Time                   `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                 gorm.DeletedAt              `gorm:"index"`
	SubjectID                 string                      `json:"SubjectID" form:"SubjectID"`
	TestType                  string                      `json:"TestType" form:"TestType"`
	Title                     string                      `json:"Title" form:"Title"`
	MeetingDate               string                      `json:"MeetingDate" form:"MeetingDate"`
	Open                      string                      `json:"Open" form:"Open"`
	Description               string                      `json:"Description" form:"Description"`
	Time                      string                      `json:"Time" form:"Time"`
	Point                     int                         `json:"Point" form:"Point"`
	Status                    string                      `json:"Status" form:"Status"`
	Attempt                   int                         `json:"Attempt" form:"Attempt"`
	GradeQuizzes              GradeQuizzes                `gorm:"foreignKey:TestTypeQuizID"`
	QuestionQuizzes           []QuestionQuizzes           `gorm:"foreignKey:TestTypeQuizID"`
	UserTestSubmissionQuizzes []UserTestSubmissionQuizzes `gorm:"foreignKey:TestTypeQuizID"`
}

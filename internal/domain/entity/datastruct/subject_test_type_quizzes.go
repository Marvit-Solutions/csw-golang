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
	GradeQuizzes              GradeQuizzes                `gorm:"foreignKey:TestTypeQuizID"`
	QuestionQuizzes           []QuestionQuizzes           `gorm:"foreignKey:TestTypeQuizID"`
	UserTestSubmissionQuizzes []UserTestSubmissionQuizzes `gorm:"foreignKey:TestTypeQuizID"`
}

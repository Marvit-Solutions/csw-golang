package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserTestSubmissionQuizzes struct {
	ID                         string                       `gorm:"type:text;primaryKey"`
	CreatedAt                  time.Time                    `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                  time.Time                    `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                  gorm.DeletedAt               `gorm:"index"`
	UserID                     string                       `json:"UserID" form:"UserID"`
	TestTypeQuizID             string                       `json:"TestTypeQuizID" form:"TestTypeQuizID"`
	Status                     string                       `json:"Status" form:"Status" gorm:"default: Selesai"` // Selesai atau Belum Dikerjakan
	SubmissionTime             time.Time                    `gorm:"default:CURRENT_TIMESTAMP" json:"SubmissionTIme" form:"SubmissionTIme"`
	GradeQuizzes               GradeQuizzes                 `gorm:"foreignKey:UserTestSubmissionQuizID"`
	UserSubmittedAnswerQuizzes []UserSubmittedAnswerQuizzes `gorm:"foreignKey:UserTestSubmissionQuizID"`
}

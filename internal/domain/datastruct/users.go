package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID                        string                      `gorm:"type:text;primaryKey"`
	CreatedAt                 time.Time                   `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt                 time.Time                   `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt                 gorm.DeletedAt              `gorm:"index"`
	RoleID                    string                      `gorm:"type:text;not null"`
	Email                     string                      `gorm:"type:varchar(64);not null;unique"`
	GoogleID                  string                      `gorm:"type:text"`
	FacebookID                string                      `gorm:"type:text"`
	Password                  string                      `gorm:"type:text;not null"`
	UserDetails               UserDetails                 `gorm:"foreignKey:UserID"`
	UserTestSubmissionQuizzes []UserTestSubmissionQuizzes `gorm:"foreignKey:UserID"`
	GradeQuizzes              GradeQuizzes                `gorm:"foreignKey:UserID"`
	UserTestimonials          UserTestimonials            `gorm:"foreignKey:UserID"`
}

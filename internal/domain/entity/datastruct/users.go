package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID                  string         `gorm:"type:text;primaryKey"`
	CreatedAt           time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	RoleID              string
	Email               string                `json:"Email" form:"Email" validate:"required,email"`
	GoogleID            string                `json:"GoogleID" form:"GoogleID"`
	FacebookID          string                `json:"FacebookID" form:"FacebookID"`
	Password            string                `json:"Password" form:"Password" validate:"required,min=8"`
	UserDetails         UserDetails           `gorm:"foreignKey:UserID"`
	UserTestSubmissions []UserTestSubmissions `gorm:"foreignKey:UserID"`
	Grades              Grades                `gorm:"foreignKey:UserID"`
}

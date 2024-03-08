package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserDetails struct {
	ID             string         `gorm:"type:text;primaryKey"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UserID         string         `gorm:"type:text;not null"`
	Name           string         `gorm:"type:varchar(100);not null"`
	PhoneNumber    string         `gorm:"type:varchar(50);not null"`
	ProfilePicture string         `gorm:"type:text;not null;default: 'assets/img/users/profile/account.png'"`
	Class          string         `gorm:"type:varchar(10);not null"`
	Addresses      Addresses      `gorm:"foreignKey:UserDetailID"`
}

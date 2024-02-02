package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserDetails struct {
	ID             string `gorm:"type:text;primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UserID         string         `json:"UserID" form:"UserID"`
	Name           string         `json:"Name" form:"Name" validate:"required"`
	PhoneNumber    string         `json:"PhoneNumber" form:"PhoneNumber" validate:"required,min=10,max=13,numeric"`
	ProfilePicture string         `json:"ProfilePicture" form:"ProfilePicture" gorm:"default: 'assets/img/users/profile/account.png'"`
	Addresses      Addresses      `gorm:"foreignKey:UserDetailID"`
	Status         string         `json:"Status" form:"Status" validate:"required"`
	TestimonialID  *string        `json:"TestimonialID" form:"TestimonialID"`
}

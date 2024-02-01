package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID             string `gorm:"type:text;primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UserID         string         `json:"UserID" form:"UserID"`
	Name           string         `json:"Name" form:"Name" validate:"required"`
	PhoneNumber    string         `json:"PhoneNumber" form:"PhoneNumber" validate:"required,min=10,max=13,numeric"`
	ProfilePicture string         `json:"ProfilePicture" form:"ProfilePicture" gorm:"default: 'assets/img/account.png'"`
	Address        Address        `gorm:"foreignKey:UserDetailID"`
}

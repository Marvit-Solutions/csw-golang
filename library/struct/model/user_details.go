package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserDetail = "user_details"

// UserDetail mapped from table <user_details>
type UserDetail struct {
	ID             int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID           string         `gorm:"column:uuid;not null" json:"uuid"`
	ClassUserID    int          `gorm:"column:class_user_id;not null" json:"class_user_id"`
	UserID         int          `gorm:"column:user_id;not null" json:"user_id"`
	Name           string         `gorm:"column:name;not null" json:"name"`
	Province       string         `gorm:"column:province;not null" json:"province"`
	Regency        string         `gorm:"column:regency;not null" json:"regency"`
	District       string         `gorm:"column:district;not null" json:"district"`
	PhoneNumber    string         `gorm:"column:phone_number;not null" json:"phone_number"`
	ProfilePicture string         `gorm:"column:profile_picture;not null" json:"profile_picture"`
	CreatedAt      time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName UserDetail's table name
func (*UserDetail) TableName() string {
	return TableNameUserDetail
}

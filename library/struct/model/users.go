package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID         int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID       string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	RoleID     int          `gorm:"column:role_id;not null" json:"role_id"`
	GoogleID   *int         `gorm:"column:google_id" json:"google_id"`
	FacebookID *int         `gorm:"column:facebook_id" json:"facebook_id"`
	Email      string         `gorm:"column:email;not null" json:"email"`
	Password   string         `gorm:"column:password;not null" json:"password"`
	CreatedAt  time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

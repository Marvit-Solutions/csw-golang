package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID         string         `gorm:"column:id;primaryKey" json:"id"`
	RoleID     string         `gorm:"column:role_id;not null" json:"role_id"`
	GoogleID   *string        `gorm:"column:google_id" json:"google_id"`
	FacebookID *string        `gorm:"column:facebook_id" json:"facebook_id"`
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

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserTestimonial = "user_testimonials"

// UserTestimonial mapped from table <user_testimonials>
type UserTestimonial struct {
	ID        int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID      string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	UserID    int          `gorm:"column:user_id;not null" json:"user_id"`
	Comment   string         `gorm:"column:comment;not null" json:"comment"`
	Rating    float64        `gorm:"column:rating;not null" json:"rating"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName UserTestimonial's table name
func (*UserTestimonial) TableName() string {
	return TableNameUserTestimonial
}

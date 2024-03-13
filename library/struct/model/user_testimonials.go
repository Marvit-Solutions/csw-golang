package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserTestimonial = "user_testimonials"

// UserTestimonial mapped from table <user_testimonials>
type UserTestimonial struct {
	ID            string         `gorm:"column:id;primaryKey" json:"id"`
	UserID        string         `gorm:"column:user_id;not null" json:"user_id"`
	TestimonialID string         `gorm:"column:testimonial_id;not null" json:"testimonial_id"`
	CreatedAt     time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName UserTestimonial's table name
func (*UserTestimonial) TableName() string {
	return TableNameUserTestimonial
}

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUserMentorTestimonial = "user_mentor_testimonials"

// UserMentorTestimonial mapped from table <user_mentor_testimonials>
type UserMentorTestimonial struct {
	ID        int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID      string         `gorm:"column:uuid;not null" json:"uuid"`
	UserID    int          `gorm:"column:user_id;not null" json:"user_id"`
	MentorID  int          `gorm:"column:mentor_id;not null" json:"mentor_id"`
	Comment   string         `gorm:"column:comment;not null" json:"comment"`
	Rating    float64        `gorm:"column:rating;not null" json:"rating"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName UserMentorTestimonial's table name
func (*UserMentorTestimonial) TableName() string {
	return TableNameUserMentorTestimonial
}

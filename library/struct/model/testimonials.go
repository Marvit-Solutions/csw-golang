package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTestimonial = "testimonials"

// Testimonial mapped from table <testimonials>
type Testimonial struct {
	ID        int          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UUID      string         `gorm:"column:uuid;not null;default:uuid_generate_v4()" json:"uuid"`
	Comment   string         `gorm:"column:comment;not null" json:"comment"`
	Rating    float64        `gorm:"column:rating;not null" json:"rating"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
}

// TableName Testimonial's table name
func (*Testimonial) TableName() string {
	return TableNameTestimonial
}

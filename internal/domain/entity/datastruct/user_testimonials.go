package datastruct

type UserTestimonials struct {
	ID            string `gorm:"type:text;primaryKey"`
	UserID        string `gorm:"type:text;not null"`
	TestimonialID string `gorm:"type:text;not null"`
}

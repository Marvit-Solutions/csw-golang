package datastruct

type Testimonials struct {
	ID         string      `json:"ID"`
	Comment    string      `json:"comment"`
	Rating     float32     `json:"rating"`
	UserDetail UserDetails `gorm:"foreignKey:TestimonialID"`
}

package datastruct

type Testimonials struct {
	ID               string           `json:"ID"`
	Comment          string           `json:"Comment"`
	Rating           float32          `json:"Rating"`
	UserTestimonials UserTestimonials `gorm:"foreignKey:TestimonialID"`
}

package datastruct

type Testimonials struct {
	ID              string           `json:"ID"`
	Comment         string           `json:"Comment"`
	Rating          float32          `json:"Rating"`
	UserTestimonial UserTestimonials `gorm:"foreignKey:TestimonialID"`
}

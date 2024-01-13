package datastruct

type Testimonials struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Status       string  `json:"status"`
	ProfilePhoto string  `json:"profile_photo"`
	Comment      string  `json:"comment"`
	Rating       float32 `json:"rating"`
}

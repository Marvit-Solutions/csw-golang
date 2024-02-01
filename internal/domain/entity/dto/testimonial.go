package dto

type ListTestimonials []struct {
	ID             string  `json:"ID"`
	Name           string  `json:"name"`
	Status         string  `json:"status"`
	ProfilePicture string  `json:"profile_photo"`
	Comment        string  `json:"comment"`
	Rating         float32 `json:"rating"`
}

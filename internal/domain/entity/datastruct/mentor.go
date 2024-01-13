package datastruct

type Mentor struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	ProfilePhoto string  `json:"profile_photo"`
	Rating       float32 `json:"rating"`
}

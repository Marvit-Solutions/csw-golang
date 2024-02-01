package datastruct

type Mentors struct {
	ID             string  `json:"ID"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	ProfilePicture string  `json:"profile_photo"`
	Rating         float32 `json:"rating"`
}

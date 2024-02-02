package dto

type ListMentor []struct {
	ID             string  `json:"ID"`
	Name           string  `json:"Name"`
	Description    string  `json:"Description"`
	ProfilePicture string  `json:"ProfilePicture"`
	Rating         float32 `json:"Rating"`
}
